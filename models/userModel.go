package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`

	ID        string    `gorm:"primary_key;unique" json:"id"`
	Name      string    `gorm:"size:100;unique;not null" json:"name"`
	Email     string    `gorm:"size:255;unique;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
