package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null;unique" json:"name"`
	Password string    `gorm:"type:varchar(255);not null" json:"password"`
}
