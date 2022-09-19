package models

import "time"

type User struct {
	UserID      string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"user_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Address     string    `json:"address,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	BirthDate   string    `json:"birth_date,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Users []User
