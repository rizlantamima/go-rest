package domain

import "time"

type Order struct {
	ID        uint32
	Fee       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
