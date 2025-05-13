package ports

import (
	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
)

type ProductsRepository interface {
	GetAllProducts(activeOnly bool) ([]domain.Product, error)
	GetProductByID(id int32) (domain.Product, error)
	CreateProduct(product domain.Product) error
	UpdateProduct(product domain.Product) error
	DeleteProductByID(id int32) error
}
