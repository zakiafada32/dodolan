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
	userController "github.com/zakiafada32/retail/api/v1/user"
	userService "github.com/zakiafada32/retail/business/user"
	"github.com/zakiafada32/retail/config"
	"github.com/zakiafada32/retail/modules/migration"
	userRepository "github.com/zakiafada32/retail/modules/user"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	// Setup
	db := config.ConnectPostgreSQL()
	migration.Migrate(db)

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(userService)

	e := echo.New()
	api.RegisterPath(e, userController)

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
