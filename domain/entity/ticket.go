package entity

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID                uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	AssetID           uuid.UUID  `json:"assetId" gorm:"type:uuid;not null"`
	Category          string     `json:"category" gorm:"not null"`
	Severity          string     `json:"severity" gorm:"check:severity IN ('low', 'medium', 'high', 'critical')"`
	Duration          int        `json:"duration"`
	DueDate           time.Time  `json:"dueDate"`
	Reporting         uuid.UUID  `json:"reporting" gorm:"type:uuid;not null"`
	AssignedTo        *uuid.UUID `json:"assignedTo" gorm:"type:uuid"`
	Comment           string     `json:"comment"`
	Status            string     `json:"status" gorm:"default:'open';check:status IN ('open', 'in_progress', 'resolved', 'closed')"`
	ResolutionComment string     `json:"resolutionComment"`
	CreatedAt         time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}