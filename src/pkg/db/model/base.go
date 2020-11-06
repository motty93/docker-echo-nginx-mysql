package model

import (
	"time"

	"gorm.io/gorm"
)

// Base redefine gorm.Model struct
type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"type:datetime(6)" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:datetime(6)" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:datetime(6)" json:"deleted_at"`
}

type GormDB struct {
	DB *gorm.DB
}
