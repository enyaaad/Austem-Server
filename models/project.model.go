package models

import "github.com/google/uuid"

type Project struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"`
	Name    string    `gorm:"type:varchar(255);not null;unique" json:"name"`
	Project string    `gorm:"type:text" json:"project"`
}
