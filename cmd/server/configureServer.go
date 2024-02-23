package server

import (
	"fmt"
	"realtime-exchange-rates/api/handlers"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Application struct {
	server *fiber.App
	startServer func()
}


func configureApp() *Application {
  app := fiber.New()

	// app configurations
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(corsConfiguration)

	// add routes.
	appRouter := app.Group("/api/v1")
	appRouter.Post("/rates", handlers.HandleExchangeRequest)

	return &Application{
		server: app,
		startServer: startServer(app),
	}
}


func startServer(app *fiber.App) func()  {
  return func ()  {
		port := 9595
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
	ctx.Next()
	return nil;
}