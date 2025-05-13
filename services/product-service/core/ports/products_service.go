package ports

import (
	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
)

type CreateProductCommand struct {
	Name  string  `json:"name" binding:"required"`
	Size  int16   `json:"size" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

type UpdateProductCommand struct {
	ID          int32    `json:"id" binding:"required"`
	NameOpt     *string  `json:"name"`
	SizeOpt     *int16   `json:"size"`
	PriceOpt    *float32 `json:"price"`
	IsActiveOpt *bool    `json:"is_active"`
}

type ProductsService interface {
	QueryAllProducts(activeOnly bool) ([]domain.BottledWater, error)
	QueryProductByID(id int32) (domain.BottledWater, error)
	CreateProduct(createProductCommand CreateProductCommand) error
	UpdateProduct(updateProductCommand UpdateProductCommand) error
	DeleteProductByID(id int32) error
}
