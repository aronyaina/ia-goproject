package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Prompt struct {
	gorm.Model `json:"-"`

	ID        string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Tag       string    `gorm:"size:255;unique;not null" json:"tag"`
	Results   string    `gorm:"type:text" json:"results"`
	UserID    string    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (p *Prompt) BeforeSave(tx *gorm.DB) (err error) {
	if p.Tag == "" {
		return errors.New("tag is required")
	}
	if p.UserID == "" {
		return errors.New("user ID is required")
	}
	return nil
}
