package domain

import "time"

type Admin struct {
	ID        uint32
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
