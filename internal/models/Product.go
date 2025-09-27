package models

import "time"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Stock       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
