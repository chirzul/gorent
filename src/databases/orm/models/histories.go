package models

import "time"

type History struct {
	HistoryID     uint      `gorm:"primaryKey" json:"history_id,omitempty"`
	UserID        string    `json:"user_id,omitempty"`
	User          User      `json:"user_data"`
	VehicleID     uint      `json:"vehicle_id,omitempty"`
	Vehicle       Vehicle   `json:"vehicle_data"`
	StartRent     string    `json:"start_rent,omitempty"`
	EndRent       string    `json:"end_rent,omitempty"`
	Prepayment    uint      `json:"prepayment,omitempty"`
	PaymentStatus string    `json:"payment_status,omitempty"`
	ReturnStatus  string    `json:"return_status,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type Histories []History
