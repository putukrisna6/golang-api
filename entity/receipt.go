package entity

import "time"

type Receipt struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Amount    uint      `json:"amount"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
