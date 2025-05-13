package services

import (
	"time"

	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
	"github.com/zhunismp/nanow4t3r/services/product/core/errors"
	"github.com/zhunismp/nanow4t3r/services/product/core/helpers"
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
		return nil, errors.New(errors.Internal, "Failed to get all products", err)
	}
	return products, nil
}

func (s *ProductsServiceImpl) QueryProductByID(id int32) (domain.Product, error) {
	product, err := s.productsRepository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, errors.New(errors.NotFound, "Product not found", err)
	}
	return product, nil
}

func (s *ProductsServiceImpl) CreateProduct(createProductCommand ports.CreateProductCommand) error {

	if err := helpers.ValidateCreateProductCommand(createProductCommand); err != nil {
		return errors.New(errors.Validation, "Invalid create product request", err)
	}

	product := domain.Product{
		Name:     createProductCommand.Name,
		Size:     createProductCommand.Size,
		Price:    createProductCommand.Price,
		IsActive: true,
	}

	if err := s.productsRepository.CreateProduct(product); err != nil {
		return errors.New(errors.Internal, "Failed to create product", err)
	}

	return nil
}

func (s *ProductsServiceImpl) UpdateProduct(updateProductCommand ports.UpdateProductCommand) error {

	if err := helpers.ValidateUpdateProductCommand(updateProductCommand); err != nil {
		return errors.New(errors.Validation, "Invalid update product request", err)
	}

	product, err := s.productsRepository.GetProductByID(updateProductCommand.ID)
	if err != nil {
		return errors.New(errors.NotFound, "Product not found", err)
	}

	updatedProduct := domain.Product{
		ID:        product.ID,
		Name:      helpers.WithFallback(updateProductCommand.NameOpt, product.Name),
		Size:      helpers.WithFallback(updateProductCommand.SizeOpt, product.Size),
		Price:     helpers.WithFallback(updateProductCommand.PriceOpt, product.Price),
		IsActive:  helpers.WithFallback(updateProductCommand.IsActiveOpt, product.IsActive),
		UpdatedAt: time.Now(),
		CreatedAt: product.CreatedAt,
	}

	if err := s.productsRepository.UpdateProduct(updatedProduct); err != nil {
		return errors.New(errors.Internal, "Failed to update product", err)
	}

	return nil
}

func (s *ProductsServiceImpl) DeleteProductByID(id int32) error {
	err := s.productsRepository.DeleteProductByID(id)
	if err != nil {
		return errors.New(errors.NotFound, "Product not found", err)
	}
	return nil
}
