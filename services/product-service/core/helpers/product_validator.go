package helpers

import (
	"fmt"

	"github.com/zhunismp/nanow4t3r/services/product/core/ports"
)

func ValidateCreateProductCommand(createProductCommand ports.CreateProductCommand) error {
	if createProductCommand.Name == "" {
		return fmt.Errorf("name is required")
	}
	if createProductCommand.Size <= 0 {
		return fmt.Errorf("size must be greater than 0")
	}
	if createProductCommand.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	return nil
}

func ValidateUpdateProductCommand(updateProductCommand ports.UpdateProductCommand) error {
	if updateProductCommand.ID <= 0 {
		return fmt.Errorf("id must be greater than 0")
	}
	if updateProductCommand.NameOpt != nil && *updateProductCommand.NameOpt == "" {
		return fmt.Errorf("name should not be empty")
	}
	if updateProductCommand.SizeOpt != nil && *updateProductCommand.SizeOpt <= 0 {
		return fmt.Errorf("size must be greater than 0")
	}
	if updateProductCommand.PriceOpt != nil && *updateProductCommand.PriceOpt <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	return nil
}
