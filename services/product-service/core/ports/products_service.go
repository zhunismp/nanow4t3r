package ports

import (
	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
)

type ProductsService interface {
	QueryAllProducts(activeOnly bool) ([]domain.Product, error)
	QueryProductByID(id string) (domain.Product, error)
	CreateProduct(product domain.Product) error
	UpdateProduct(product domain.Product) error
	DeleteProductByID(id string) error
}
