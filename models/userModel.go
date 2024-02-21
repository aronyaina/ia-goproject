package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`

	UserID    string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserName  string    `gorm:"size:100;unique;not null" json:"name"`
	UserMail  string    `gorm:"size:255;unique;not null" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
