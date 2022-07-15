package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Service struct {
	gorm.Model
	Name    string  `json:"name"`
	Purpose string  `json:"purpose"`
	SLA     float64 `json:"sla"`
	Price   float64 `json:"price"`
}
