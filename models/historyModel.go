package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model `json:"-"`

	ID        string    `gorm:"primary_key;unique" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	UserID    string    `gorm:"not null" json:"userId"`
	PromptID  string    `gorm:"not null" json:"promptId"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	User   *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Prompt *Prompt `gorm:"foreignKey:PromptID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (h *History) BeforeSave(tx *gorm.DB) (err error) {
	if h.Title == "" {
		return errors.New("title is required")
	}
	if h.UserID == "" {
		return errors.New("user ID is required")
	}
	if h.PromptID == "" {
		return errors.New("prompt ID is required")
	}
	return nil
}
