package web

import (
	"app/internal/store"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, namesStore *store.NamesStore) {
	app.Get("/", GetMain(namesStore))
	app.Post("/generate", PostGenerate(namesStore))
}
