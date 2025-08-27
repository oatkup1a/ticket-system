package entity

import "time"

type Application struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name        string    `gorm:"size:255;not null"     json:"name"`
	Description string    `gorm:"type:text"             json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
