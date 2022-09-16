package models

import "time"

type User struct {
	UserID      string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"user_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	DisplayName string    `json:"display_name"`
	BirthDate   string    `json:"birth_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Users []User
