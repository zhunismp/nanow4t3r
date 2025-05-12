package dtos

type CreateProductRequest struct {
	Name      string    `json:"name"`
	Size      int16    `json:"size"`
	Price     float32   `json:"price"`
}