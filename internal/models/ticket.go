package models

import "time"

type Ticket struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Allocation  int       `json:"allocation"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Ticket) TableName() string {
	return "tickets"
}
