package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-backend/graph/generated"
	"homework-backend/graph/model"
	"homework-backend/storage"
	"homework-backend/utils/validator"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateAnswers(ctx context.Context, input []*model.NewAnswer) (*model.Result, error) {
	hasErrors := false

	dbQuestions, err := storage.ListAllQuestions(r.Sqlite)
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	// start - make sure all questions are answered
	var dbQuestionIDs []string

	for _, dbQuestion := range dbQuestions {
		dbQuestionIDs = append(dbQuestionIDs, dbQuestion.ID)
	}

	questionIDsMap := make(map[string]bool)

	for _, answer := range input {
		questionIDsMap[answer.QuestionID] = true
	}

	var missingQuestionIds []string

	for _, dbQuestionID := range dbQuestionIDs {
		if !questionIDsMap[dbQuestionID] {
			missingQuestionIds = append(missingQuestionIds, dbQuestionID)
		}
	}

	if len(missingQuestionIds) > 0 {
		hasErrors = true
		graphql.AddError(ctx, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"type":        "missingQuestion",
				"questionIDs": missingQuestionIds,
			},
			Message: "You need to answer all questions.",
			Path:    graphql.GetPath(ctx),
		})
	}
	// end - make sure all questions are answered

	mapDbQuestions := storage.MapQuestions(dbQuestions)
	var dbAnswers []*storage.Answer

	for _, answer := range input {
		dbQuestion := mapDbQuestions[answer.QuestionID]

		err = validator.ValidateAnswer(answer, mapDbQuestions[answer.QuestionID])
		if err != nil {
			hasErrors = true
			graphql.AddError(ctx, &gqlerror.Error{
				Extensions: map[string]interface{}{
					"type":       "question",
					"questionID": answer.QuestionID,
				},
				Message: err.Error(),
				Path:    graphql.GetPath(ctx),
			})

			continue
		}

		var weight = dbQuestion.Weight

		if dbQuestion.Type == storage.QuestionTypeChoice {
			mapOptions := storage.MapOptions(dbQuestion.Options)

			weight = mapOptions[*answer.OptionID].Weight
		}

		dbAnswers = append(dbAnswers, &storage.Answer{
			Body:       answer.Body,
			OptionID:   answer.OptionID,
			QuestionID: answer.QuestionID,
			Weight:     weight,
		})
	}

	if hasErrors {
		return nil, nil
	}

	err = storage.CreateAnswers(r.Sqlite, dbAnswers)
	if err != nil {
		return nil, gqlerror.Errorf("An error has occurred (but we are on it).")
	}

	var weightSum float64 = 0

	for _, dbAnswer := range dbAnswers {
		weightSum += dbAnswer.Weight
	}

	result := &model.Result{
		Weight: weightSum / float64(len(dbAnswers)),
	}

	j, err := json.Marshal(dbAnswers)
	if err != nil {
		log.Println("An error has occurred while printing the answers to stdout.")
	} else {
		fmt.Println(string(j))
	}

	return result, nil
}

func (r *queryResolver) Questions(ctx context.Context) ([]model.Question, error) {
	questions := make([]model.Question, 0)

	var dbQuestions []storage.Question

	err := r.Sqlite.
		Model(&storage.Question{}).
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
					ID:     dbOption.ID,
					Body:   dbOption.Body,
					Weight: dbOption.Weight,
				})
			}

			questions = append(questions, model.ChoiceQuestion{
				ID:      dbQuestion.ID,
				Body:    dbQuestion.Body,
				Weight:  dbQuestion.Weight,
				Options: options,
			})
		} else if dbQuestion.Type == storage.QuestionTypeText {
			questions = append(questions, model.TextQuestion{
				ID:     dbQuestion.ID,
				Body:   dbQuestion.Body,
				Weight: dbQuestion.Weight,
			})
		}
	}

	return questions, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
