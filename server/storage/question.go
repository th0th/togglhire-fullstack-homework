package storage

import (
	"gorm.io/gorm"
	"homework-backend/utils/uuid"
)

const QuestionTypeChoice = "CHOICE"
const QuestionTypeText = "TEXT"

type Question struct {
	ID string `gorm:"primaryKey"`

	Type string

	Body   string  `gorm:"not null"`
	Weight float64 `gorm:"not null"`

	Options []*Option

	Answer *Answer
}

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == "" {
		q.ID = uuid.GenerateUuid()
	}

	return
}

func (q *Question) PopulateOptions(tx *gorm.DB) (err error) {
	err = tx.Model(&Option{}).Where("question_id = ?", q.ID).Find(&q.Options).Error

	return
}

func ListAllQuestions(tx *gorm.DB) (questions []*Question, err error) {
	err = tx.Model(&Question{}).Preload("Options").Find(&questions).Error

	return
}

func MapQuestions(questions []*Question) map[string]*Question {
	m := make(map[string]*Question)

	for _, question := range questions {
		m[question.ID] = question
	}

	return m
}
