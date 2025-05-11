package infrastructures

import (
	"github.com/gin-gonic/gin"
	"github.com/zhunismp/nanow4t3r/services/product/adapters/handlers"
	"github.com/zhunismp/nanow4t3r/services/product/adapters/repositories"
	"github.com/zhunismp/nanow4t3r/services/product/core/services"
	"github.com/zhunismp/nanow4t3r/services/product/infrastructures/config"
	"github.com/zhunismp/nanow4t3r/services/product/infrastructures/db"
)

func Start() {
	cfg := config.LoadConfig()

	dbClient := db.GetNewGormDBInstance(cfg.DB_CONFIG)

	productRepository := repositories.NewProductsRepositoryImpl(dbClient)
	productService := services.NewProductsServiceImpl(productRepository)
	productHttpHandler := handlers.NewProductHttpHandler(productService)

	app := gin.Default()
	route(app, productHttpHandler)
	app.Run(":" + cfg.APP_CONFIG.PORT)
}

func route(app *gin.Engine, productHttpHandler *handlers.ProductHttpHandler) {
	app.GET("/products", productHttpHandler.GetAllProducts)
	app.GET("/products/:id", productHttpHandler.GetProductByID)
}
