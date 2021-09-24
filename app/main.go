package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zakiafada32/retail/api"
	categoryController "github.com/zakiafada32/retail/api/v1/category"
	productController "github.com/zakiafada32/retail/api/v1/product"
	userController "github.com/zakiafada32/retail/api/v1/user"
	categoryService "github.com/zakiafada32/retail/business/category"
	productService "github.com/zakiafada32/retail/business/product"
	userService "github.com/zakiafada32/retail/business/user"
	"github.com/zakiafada32/retail/config"
	"github.com/zakiafada32/retail/modules"
	categoryRepository "github.com/zakiafada32/retail/modules/category"
	productRepository "github.com/zakiafada32/retail/modules/product"
	userRepository "github.com/zakiafada32/retail/modules/user"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	// Setup
	db := config.ConnectPostgreSQL()
	modules.Migrate(db)

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(userService)

	productRepository := productRepository.NewProductRepository(db)
	productService := productService.NewProductService(productRepository)
	productController := productController.NewProductController(productService)

	categoryRepository := categoryRepository.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepository)
	categoryController := categoryController.NewCategoryController(categoryService)

	e := echo.New()
	api.Bootstrap(e, userController, categoryController, productController)

	// Start server
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	go func() {
		if err := e.Start(":3000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
