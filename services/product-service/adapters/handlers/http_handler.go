package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
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
	activeOnly := c.Query("active_only") == "true"
	products, err := s.productsService.QueryAllProducts(activeOnly)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, products)
}

func (s *ProductHttpHandler) GetProductByID(c *gin.Context) {
	idOpt := c.Param("id")
	id, err := strconv.ParseUint(idOpt, 10, 32)
    if err != nil {
        c.JSON(404, gin.H{"error": "Product not found"})
		return
    }

	product, err := s.productsService.QueryProductByID(uint32(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}
