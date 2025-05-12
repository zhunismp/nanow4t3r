package infrastructures

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	runWithGracefulShutdown(app, cfg.APP_CONFIG.PORT)
}

func runWithGracefulShutdown(app *gin.Engine, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown: ", err)
	}

	<-ctx.Done()

	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting...")
}

func route(app *gin.Engine, productHttpHandler *handlers.ProductHttpHandler) {
	// Health check endpoint
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Product service is running 🐳"})
	})

	// Product endpoints
	app.GET("/products", productHttpHandler.GetAllProducts)
	app.GET("/products/:id", productHttpHandler.GetProductByID)
	app.POST("/products", productHttpHandler.CreateProduct)
	app.PUT("/products", productHttpHandler.UpdateProduct)
}
