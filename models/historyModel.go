package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model `json:"-"`

	ID        string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	UserID    string    `gorm:"not null" json:"user_id"`
	PromptID  string    `gorm:"not null" json:"prompt_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

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
