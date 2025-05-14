package ports

import (
	"github.com/zhunismp/nanow4t3r/services/product/adapters/dtos"
	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
)

type ProductsService interface {
	QueryAllProducts(activeOnly bool) ([]domain.BottledWater, error)
	QueryProductByID(id int32) (domain.BottledWater, error)
	CreateProduct(createProductCommand dtos.CreateProductRequest) error
	UpdateProduct(updateProductCommand dtos.UpdateProductRequest) error
	DeleteProductByID(id int32) error
}
