package models

import "github.com/google/uuid"

type Product struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null;" json:"name"`
	Count      int       `json:"count"`
	Cost       int       `json:"cost"`
	IsSelected bool      `gorm:"default:false;" json:"isSelected"`
}
