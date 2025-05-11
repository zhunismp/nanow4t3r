package domain

import "time"

type Product struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Size      uint16    `json:"size"`  // in CC
	Price     float32   `json:"price"` // in Bath
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
