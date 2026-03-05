package web

// middleware to create id cookie that sets a uuid if not present
import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func IdLocalsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Cookies("id")
		newId := uuid.New().String()
		if id == "" {
			c.Cookie(&fiber.Cookie{
				Name:     "id",
				Path:     "/",
				SameSite: "Lax",
				HTTPOnly: true,
				Value:    newId,
			})
			c.Locals("id", newId)
			return c.Next()
		}
		c.Locals("id", id)
		return c.Next()
	}
}

func RegisterMiddleware(app *fiber.App) {
	app.Use(IdLocalsMiddleware())
}
