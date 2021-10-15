package storage

import (
	"gorm.io/gorm"
	"homework-backend/utils/uuid"
)

type Answer struct {
	ID string `gorm:"primaryKey"`

	Question   Question `gorm:"constraint:OnDelete:CASCADE"`
	QuestionID string   `gorm:"not null"`

	Body *string

	Option   *Option `gorm:"constraint:OnDelete:CASCADE"`
	OptionID *string

	Weight float64 `gorm:"not null"`
}

func (a *Answer) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == "" {
		a.ID = uuid.GenerateUuid()
	}

	return
}

func CreateAnswers(tx *gorm.DB, answers []*Answer) (err error) {
	err = tx.Model(&Answer{}).Create(&answers).Error

	return
}
