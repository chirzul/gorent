package models

import "time"

type Vehicle struct {
	VehicleID   uint      `gorm:"primaryKey" json:"vehicle_id,omitempty"`
	Name        string    `gorm:"not null" json:"name,omitempty"`
	Location    string    `gorm:"not null" json:"location,omitempty"`
	Description string    `gorm:"not null" json:"description,omitempty"`
	Price       uint      `gorm:"not null" json:"price,omitempty"`
	Status      string    `gorm:"not null" json:"status,omitempty"`
	Stock       uint      `gorm:"not null" json:"stock,omitempty"`
	Category    string    `gorm:"not null" json:"category,omitempty"`
	Image       string    `gorm:"not null" json:"image,omitempty"`
	TotalRented int       `gorm:"not null" json:"total_rented"`
	CreatedAt   time.Time `gorm:"default:now(); not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"default:now(); not null" json:"updated_at,omitempty"`
}

type Vehicles []Vehicle
