package models

import "time"

type History struct {
	HistoryID     uint      `gorm:"primaryKey" json:"history_id"`
	UserID        string    `json:"user_id"`
	VehicleID     uint      `json:"vehicle_id"`
	StartRent     string    `json:"start_rent"`
	EndRent       uint      `json:"end_rent"`
	Prepayment    string    `json:"prepayment"`
	PaymentStatus uint      `json:"payment_status"`
	ReturnStatus  string    `json:"return_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Histories []History
