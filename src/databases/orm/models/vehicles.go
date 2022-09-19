package models

import "time"

type Vehicle struct {
	VehicleID   uint      `gorm:"primaryKey" json:"vehicle_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Location    string    `json:"location,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       uint      `json:"price,omitempty"`
	Status      string    `json:"status,omitempty"`
	Stock       uint      `json:"stock,omitempty"`
	Category    string    `json:"category,omitempty"`
	Image       string    `json:"image,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	TotalRented int       `json:"total_rented"`
}

type Vehicles []Vehicle
