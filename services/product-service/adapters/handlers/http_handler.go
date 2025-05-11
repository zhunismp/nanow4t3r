package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhunismp/nanow4t3r/services/product/core/ports"
)

type ProductHttpServer struct {
	productsService ports.ProductsService
}

func NewProductHttpServer(productsService ports.ProductsService) *ProductHttpServer {
	return &ProductHttpServer{
		productsService: productsService,
	}
}

func (s *ProductHttpServer) GetAllProducts(c *gin.Context) {
	activeOnly := c.Query("active_only") == "true"
	products, err := s.productsService.QueryAllProducts(activeOnly)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, products)
}

func (s *ProductHttpServer) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := s.productsService.QueryProductByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}
