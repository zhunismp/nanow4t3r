package services

import (
	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
	"github.com/zhunismp/nanow4t3r/services/product/core/ports"
)

type ProductsServiceImpl struct {
	productsRepository ports.ProductsRepository
}

func NewProductsServiceImpl(productsRepository ports.ProductsRepository) *ProductsServiceImpl {
	return &ProductsServiceImpl{
		productsRepository: productsRepository,
	}
}

func (s *ProductsServiceImpl) QueryAllProducts(activeOnly bool) ([]domain.Product, error) {
	products, err := s.productsRepository.GetAllProducts(activeOnly)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductsServiceImpl) QueryProductByID(id uint32) (domain.Product, error) {
	product, err := s.productsRepository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *ProductsServiceImpl) CreateProduct(product domain.Product) error {
	err := s.productsRepository.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductsServiceImpl) UpdateProduct(product domain.Product) error {
	err := s.productsRepository.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductsServiceImpl) DeleteProductByID(id uint32) error {
	err := s.productsRepository.DeleteProductByID(id)
	if err != nil {
		return err
	}
	return nil
}
