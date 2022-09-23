package models

import "time"

type User struct {
	UserID    string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"user_id,omitempty"`
	Username  string    `gorm:"not null" json:"username,omitempty" validate:"required"`
	Email     string    `gorm:"not null" json:"email,omitempty" validate:"required"`
	Role      string    `gorm:"not null" json:"role,omitempty"`
	Password  string    `gorm:"not null" json:"password,omitempty" validate:"required"`
	Name      string    `gorm:"not null" json:"name,omitempty"`
	Gender    string    `gorm:"not null" json:"gender,omitempty"`
	Address   string    `gorm:"not null" json:"address,omitempty"`
	Phone     string    `gorm:"not null" json:"phone,omitempty"`
	BirthDate string    `gorm:"not null" json:"birth_date,omitempty"`
	CreatedAt time.Time `gorm:"default:now(); not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"default:now(); not null" json:"updated_at,omitempty"`
}

type Users []User
