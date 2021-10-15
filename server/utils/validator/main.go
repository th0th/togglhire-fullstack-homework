package validator

import (
	"homework-backend/graph/model"
	"homework-backend/storage"
)

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

func ValidateAnswer(a *model.NewAnswer, q *storage.Question) error {
	if q == nil {
		return &ValidationError{
			Message: "Please provide a valid question ID.",
		}
	}

	if q.Type == storage.QuestionTypeChoice {
		e := &ValidationError{
			Message: "Invalid option.",
		}

		if a.OptionID == nil {
			return e
		}

		mapOptions := storage.MapOptions(q.Options)

		if mapOptions[*a.OptionID] == nil {
			return e
		}
	}

	if q.Type == storage.QuestionTypeText {
		e := &ValidationError{
			Message: "Please answer.",
		}

		if a.Body == nil || *a.Body == "" {
			return e
		}
	}

	return nil
}
