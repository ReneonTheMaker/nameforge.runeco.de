package web

import (
	"log"

	"app/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

type App struct {
	FiberApp   *fiber.App
	NamesStore *store.NamesStore
}

func NewApp() *App {
	log.Println("Initializing web application...")
	// Template Engine golang html/template
	engine := html.New("./views", ".html")

	registerRenderFunctions(engine)

	// Fiber App
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	// Static files
	app.Static("/static", "./static")

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Create the names store
	namesStore := store.NewNamesStore()
	renderConfigStore := store.NewRenderConfigStore()

	// Middleware to set ID cookie
	RegisterMiddleware(app)

	// Register Routes - defined in routes.go
	RegisterRoutes(app, namesStore, renderConfigStore)

	log.Println("Web application initialized successfully.")

	return &App{
		FiberApp:   app,
		NamesStore: namesStore,
	}
}
