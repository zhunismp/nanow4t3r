package dtos

import (
	"fmt"

	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
)

func MapCreateProductRequestToProduct(request CreateProductRequest) (domain.Product, error) {

	if err := validateCreateProductRequest(request); err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
		Name:     request.Name,
		Size:     uint16(request.Size),
		Price:    request.Price,
		IsActive: true,
	}, nil
}

func validateCreateProductRequest(request CreateProductRequest) error {
	if request.Name == "" {
		return fmt.Errorf("name is required")
	}
	if request.Size <= 0 {
		return fmt.Errorf("size must be greater than 0")
	}
	if request.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	return nil
}
