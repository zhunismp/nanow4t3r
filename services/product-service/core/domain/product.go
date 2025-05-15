package domain

import "time"

type BottledWater struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Size      int16     `json:"size"`  // size in CC
	Price     float32   `json:"price"` // currency in Bath
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
