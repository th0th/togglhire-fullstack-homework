package storage

import (
	"gorm.io/gorm"
	"homework-backend/utils/uuid"
)

type Result struct {
	ID string `gorm:"primaryKey"`

	Weight float64 `gorm:"not null"`
}

func (r *Result) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = uuid.GenerateUuid()
	}

	return
}
