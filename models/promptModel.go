package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Prompt struct {
	gorm.Model `json:"-"`

	ID        string    `gorm:"primary_key;unique" json:"id"`
	Tag       string    `gorm:"size:255;not null" json:"tag"`
	Result    string    `gorm:"type:text" json:"result"`
	ImageName string    `gorm:"type:text" json:"image_name"`
	UserID    string    `gorm:"not null" json:"userId"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

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
