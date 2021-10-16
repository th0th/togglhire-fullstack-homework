package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"homework-backend/graph/generated"
	"homework-backend/graph/model"
	"homework-backend/storage"
	"homework-backend/utils/validator"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) SetAnswer(ctx context.Context, input model.AnswerInput) (model.Answer, error) {
	dbQuestion := &storage.Question{}

	err := r.Sqlite.Model(&storage.Question{}).Where("id = ?", input.QuestionID).First(dbQuestion).Error
	if err != nil {
		return nil, gqlerror.Errorf("Please provide a valid questionID.")
	}

	err = dbQuestion.PopulateOptions(r.Sqlite)
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	err = validator.ValidateAnswer(&input, dbQuestion)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}

	dbAnswer := &storage.Answer{}

	_ = r.Sqlite.Model(&storage.Answer{}).Where("question_id = ?", input.QuestionID).First(&dbAnswer).Error

	dbAnswer.QuestionID = input.QuestionID
	dbAnswer.OptionID = input.OptionID
	dbAnswer.Body = input.Body
	dbAnswer.Weight = dbQuestion.Weight

	if dbQuestion.Type == storage.QuestionTypeChoice {
		mapDbOptions := storage.MapOptions(dbQuestion.Options)

		dbOption := mapDbOptions[*dbAnswer.OptionID]
		dbAnswer.Weight = dbOption.Weight
	}

	err = r.Sqlite.Save(dbAnswer).Error
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	if dbQuestion.Type == storage.QuestionTypeChoice {
		return &model.ChoiceQuestionAnswer{
			ID:         dbAnswer.ID,
			QuestionID: dbAnswer.QuestionID,
			Weight:     dbAnswer.Weight,
			OptionID:   dbAnswer.OptionID,
		}, nil
	} else if dbQuestion.Type == storage.QuestionTypeText {
		return &model.TextQuestionAnswer{
			ID:         dbAnswer.ID,
			QuestionID: dbAnswer.QuestionID,
			Weight:     dbAnswer.Weight,
			Body:       dbAnswer.Body,
		}, nil
	}

	return nil, nil
}

func (r *mutationResolver) Submit(ctx context.Context) (*model.Result, error) {
	var dbResultCount int64 = 0

	err := r.Sqlite.Model(&storage.Result{}).Count(&dbResultCount).Error
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	if dbResultCount != 0 {
		return nil, gqlerror.Errorf("You have already completed.")
	}

	var dbQuestions []*storage.Question
	var dbAnswers []*storage.Answer

	err = r.Sqlite.Model(&storage.Question{}).Find(&dbQuestions).Error
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	err = r.Sqlite.Model(&storage.Answer{}).Find(&dbAnswers).Error
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	if len(dbQuestions) != len(dbAnswers) {
		return nil, gqlerror.Errorf("You need to answer all questions.")
	}

	var weightSum float64 = 0

	for _, dbAnswer := range dbAnswers {
		weightSum += dbAnswer.Weight
	}

	dbResult := &storage.Result{
		Weight: weightSum / float64(len(dbAnswers)),
	}

	err = r.Sqlite.Save(dbResult).Error
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	result := &model.Result{
		Weight: dbResult.Weight,
	}

	return result, nil
}

func (r *queryResolver) Questions(ctx context.Context) ([]model.Question, error) {
	questions := make([]model.Question, 0)

	var dbQuestions []storage.Question

	err := r.Sqlite.
		Model(&storage.Question{}).
		Preload("Answer").
		Find(&dbQuestions).
		Error
	if err != nil {
		return nil, err
	}

	for _, dbQuestion := range dbQuestions {
		if dbQuestion.Type == storage.QuestionTypeChoice {
			options := make([]*model.Option, 0)

			err = dbQuestion.PopulateOptions(r.Sqlite)
			if err != nil {
				return nil, err
			}

			for _, dbOption := range dbQuestion.Options {
				options = append(options, &model.Option{
					ID:   dbOption.ID,
					Body: dbOption.Body,
				})
			}

			choiceQuestion := model.ChoiceQuestion{
				ID:      dbQuestion.ID,
				Body:    dbQuestion.Body,
				Weight:  dbQuestion.Weight,
				Options: options,
			}

			if dbQuestion.Answer != nil {
				choiceQuestion.Answer = &model.ChoiceQuestionAnswer{
					ID:         dbQuestion.Answer.ID,
					QuestionID: dbQuestion.Answer.QuestionID,
					Weight:     dbQuestion.Answer.Weight,
					OptionID:   dbQuestion.Answer.OptionID,
				}
			}

			questions = append(questions, choiceQuestion)
		} else if dbQuestion.Type == storage.QuestionTypeText {
			textQuestion := model.TextQuestion{
				ID:     dbQuestion.ID,
				Body:   dbQuestion.Body,
				Weight: dbQuestion.Weight,
			}

			if dbQuestion.Answer != nil {
				textQuestion.Answer = &model.TextQuestionAnswer{
					ID:         dbQuestion.Answer.ID,
					QuestionID: dbQuestion.Answer.QuestionID,
					Weight:     dbQuestion.Answer.Weight,
					Body:       dbQuestion.Answer.Body,
				}
			}

			questions = append(questions, textQuestion)
		}
	}

	return questions, nil
}

func (r *queryResolver) Result(ctx context.Context) ([]*model.Result, error) {
	var dbResults []*storage.Result

	err := r.Sqlite.Model(&storage.Result{}).Find(&dbResults).Error
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	results := make([]*model.Result, 0)

	for _, dbResult := range dbResults {
		results = append(results, &model.Result{
			Weight: dbResult.Weight,
		})
	}

	return results, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
