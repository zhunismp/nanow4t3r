package dtos

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Size  int16   `json:"size" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

type UpdateProductRequest struct {
	ID          int32    `json:"id" binding:"required"`
	NameOpt     *string  `json:"name"`
	SizeOpt     *int16   `json:"size"`
	PriceOpt    *float32 `json:"price"`
	IsActiveOpt *bool    `json:"is_active"`
}
