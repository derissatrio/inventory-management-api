package entity

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Area        string    `json:"area" gorm:"not null"`
	Description string    `json:"description"`
	Capacity    int       `json:"capacity" gorm:"default:0"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}