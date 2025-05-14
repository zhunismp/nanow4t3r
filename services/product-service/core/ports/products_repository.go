package ports

import (
	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
)

type ProductsRepository interface {
	GetAllProducts(activeOnly bool) ([]domain.BottledWater, error)
	GetProductByID(id int32) (domain.BottledWater, error)
	CreateProduct(product domain.BottledWater) error
	UpdateProduct(product domain.BottledWater) error
	DeleteProductByID(id int32) error
}
