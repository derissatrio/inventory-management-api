package entity

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID            uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UniqueID      string     `json:"uniqueId" gorm:"not null"`
	Name          string     `json:"name" gorm:"not null"`
	Comment       string     `json:"comment"`
	Detail        string     `json:"detail"`
	Qty           int        `json:"qty" gorm:"default:1"`
	Brand         string     `json:"brand"`
	Type          string     `json:"type" gorm:"check:type IN ('it', 'non_it')"`
	Status        string     `json:"status" gorm:"default:'available';check:status IN ('available', 'booked', 'broken', 'repair')"`
	Category      string     `json:"category"`
	LocationID    *uuid.UUID `json:"locationId" gorm:"type:uuid"`
	LocationLabel string     `json:"locationLabel"`
	Location      *Location  `json:"location,omitempty" gorm:"foreignKey:LocationID;references:ID"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}