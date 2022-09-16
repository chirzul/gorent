package models

import "time"

type Vehicle struct {
	VehicleID   uint      `gorm:"primaryKey" json:"vehicle_id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Price       uint      `json:"price"`
	Status      string    `json:"status"`
	Stock       uint      `json:"stock"`
	Category    string    `json:"category"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Vehicles []Vehicle
