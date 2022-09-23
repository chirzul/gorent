package models

import "time"

type History struct {
	HistoryID     uint      `gorm:"primaryKey" json:"history_id,omitempty"`
	UserID        string    `gorm:"not null" json:"user_id,omitempty"`
	User          User      `gorm:"not null" json:"user_data"`
	VehicleID     uint      `gorm:"not null" json:"vehicle_id,omitempty"`
	Vehicle       Vehicle   `gorm:"not null" json:"vehicle_data"`
	StartRent     string    `gorm:"not null" json:"start_rent,omitempty"`
	EndRent       string    `gorm:"not null" json:"end_rent,omitempty"`
	Prepayment    uint      `gorm:"not null" json:"prepayment,omitempty"`
	PaymentStatus string    `gorm:"not null" json:"payment_status,omitempty"`
	ReturnStatus  string    `gorm:"not null" json:"return_status,omitempty"`
	CreatedAt     time.Time `gorm:"default:now(); not null" json:"created_at,omitempty"`
	UpdatedAt     time.Time `gorm:"default:now(); not null" json:"updated_at,omitempty"`
}

type Histories []History
