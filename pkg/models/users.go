package models

import (
	"time"

	"gorm.io/gorm"
)

// User is a struct that represents a user
type User struct {
	ID        uint64         `gorm:"primary_key;auto_increment" json:"id"`
	Name      string         `json:"name,omitempty"`
	Nick      string         `json:"nick,omitempty"`
	Email     string         `json:"email,omitempty"`
	Password  string         `json:"password,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
