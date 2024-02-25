package server

import (
	"fmt"
	"net/http"
	handlers "realtime-exchange-rates/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Application struct {
	server *fiber.App
	startServer func()
}

type ApplicationError struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type ErrorHandlerFunc func (ctx *fiber.Ctx, err error) error

func configureApp() *Application {
  app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler(),
		AppName: "Cadana",
	})

	// app configurations
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(corsConfiguration)

	// add routes.
	setupRoute(app)
	
	return &Application{
		server: app,
		startServer: startServer(app),
	}
}

func setupRoute(app *fiber.App) {
	appRouter := app.Group("/api/v1")
	appRouter.Get("/rates", handlers.HandleRealtimeExchangeRate)
	appRouter.Get("/persons", handlers.GetPersonsInformation)
}

func startServer(app *fiber.App) func()  {
  return func ()  {
		port := 3000
		fmt.Printf("Server is running on http://localhost:%d\n", port)
		err := app.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	}
}

func corsConfiguration(ctx *fiber.Ctx) error {
	ctx.Set("Access-Control-Allow-Origin", "*")
	ctx.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Set("Access-Control-Allow-Headers", "Content-Type")
	return ctx.Next()
}

func errorHandler() ErrorHandlerFunc {
	return func (ctx *fiber.Ctx, err error) error  {
		return ctx.Status(ctx.Response().StatusCode()).JSON(ApplicationError {
		Success: false,
		Message: http.StatusText(500),
	})
	}
}