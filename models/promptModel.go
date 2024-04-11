package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Prompt struct {
	gorm.Model `json:"-"`

	ID     string `gorm:"primary_key;unique" json:"id"`
	Tag    string `gorm:"size:255;not null" json:"tag"`
	Result string `gorm:"type:text" json:"result"`
	Input  string `gorm:"type:text" json:"input"`
	//UserID    string    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (p *Prompt) BeforeSave(tx *gorm.DB) (err error) {
	if p.Tag == "" {
		return errors.New("tag is required")
	}
	return nil
}
