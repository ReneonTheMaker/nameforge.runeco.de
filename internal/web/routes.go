package web

import (
	"app/internal/store"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, namesStore *store.NamesStore, renderConfigStore *store.RenderConfigStore) {
	app.Get("/", GetMain(namesStore, renderConfigStore))
	app.Post("/generate", PostGenerate(namesStore, renderConfigStore))
}
