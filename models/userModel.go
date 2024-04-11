package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`

	ID        string    `gorm:"primary_key;unique" json:"id"`
	Email     string    `gorm:"size:255;unique;not null" json:"email"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	//Prompt []Prompt `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
