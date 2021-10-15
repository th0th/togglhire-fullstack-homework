package storage

import (
	"gorm.io/gorm"
	"homework-backend/utils/uuid"
)

type Option struct {
	ID         string   `gorm:"primaryKey"`
	Question   Question `gorm:"constraint:OnDelete:CASCADE"`
	QuestionID string   `gorm:"not null"`

	Body   string  `gorm:"not null"`
	Weight float64 `gorm:"not null"`
}

func (o *Option) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == "" {
		o.ID = uuid.GenerateUuid()
	}

	return
}

func MapOptions(options []*Option) map[string]*Option {
	m := make(map[string]*Option)

	for _, option := range options {
		m[option.ID] = option
	}

	return m
}
