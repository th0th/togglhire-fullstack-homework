package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"homework-backend/graph/generated"
	"homework-backend/graph/model"
)

func (r *queryResolver) Questions(ctx context.Context) ([]model.Question, error) {
	return []model.Question{
		model.ChoiceQuestion{
			ID:     "100",
			Body:   "Where does the sun set?",
			Weight: 0.5,
			Options: []*model.Option{
				{ID: "200", Body: "East", Weight: 0},
				{ID: "201", Body: "West", Weight: 1},
			},
		},
		model.TextQuestion{
			ID:     "101",
			Body:   "What is your favourite food?",
			Weight: 1,
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
