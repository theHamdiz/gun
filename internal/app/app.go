package app

import (
	"log"

	"github.com/theHamdiz/gun/internal/middleware"
	"github.com/theHamdiz/gun/internal/routes"
	"github.com/theHamdiz/gun/internal/utils"
)

type App struct {
	Fiber *fiber.App
}

func New() *App {
	// Initialize Fiber app
	app := &App{
		Fiber: fiber.New(),
	}

	// Setup middleware
	app.setupMiddleware()

	// Register routes
	app.registerRoutes()

	return app
}

func (a *App) setupMiddleware() {
	// Example: a.Fiber.Use(middleware.Logger())
	a.Fiber.Use(middleware.Recover())
	// Add custom middleware
	a.Fiber.Use(middleware.AuthMiddleware())
}

func (a *App) registerRoutes() {
	routes.RegisterRoutes(a.Fiber)
}

func (a *App) Run(address string) {
	// Setup graceful shutdown
	utils.SetupGracefulShutdown(a.Fiber)

	// Start the server
	if err := a.Fiber.Listen(address); err != nil {
		log.Fatal(err)
	}
}
