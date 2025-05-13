package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhunismp/nanow4t3r/services/product/core/errors"
	"github.com/zhunismp/nanow4t3r/services/product/core/ports"
)

type ProductHttpHandler struct {
	productsService ports.ProductsService
}

func NewProductHttpHandler(productsService ports.ProductsService) *ProductHttpHandler {
	return &ProductHttpHandler{
		productsService: productsService,
	}
}

func (s *ProductHttpHandler) GetAllProducts(c *gin.Context) {

	// Check if the query parameter "active_only" is present and set to "false"
	// Otherwise, default to true
	activeOnly := c.Query("active_only") != "false"
	products, err := s.productsService.QueryAllProducts(activeOnly)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(200, products)
}

func (s *ProductHttpHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	product, err := s.productsService.QueryProductByID(int32(id))
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(200, product)
}

func (s *ProductHttpHandler) CreateProduct(c *gin.Context) {
	var createProductCommand ports.CreateProductCommand
	if err := c.ShouldBindJSON(&createProductCommand); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	if err := s.productsService.CreateProduct(createProductCommand); err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(201, gin.H{"message": "Product created successfully"})
}

func (s *ProductHttpHandler) UpdateProduct(c *gin.Context) {
	var updateProductCommand ports.UpdateProductCommand
	if err := c.ShouldBindJSON(&updateProductCommand); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	if err := s.productsService.UpdateProduct(updateProductCommand); err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Product updated successfully"})
}

func (s *ProductHttpHandler) DeleteProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	if err := s.productsService.DeleteProductByID(int32(id)); err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}

func handleServiceError(c *gin.Context, err error) {
	if appErr, ok := err.(*errors.AppError); ok {
		c.Set(errors.AppErrorKey, appErr)
		return
	}

	c.Error(err)
}
