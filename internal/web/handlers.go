package web

import (
	"strings"

	"app/internal/store"

	"github.com/gofiber/fiber/v2"
)

// index page handler
func GetMain(namesStore *store.NamesStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get id cookie
		id := c.Locals("id").(string)

		// get names for id
		err := ""
		if id == "" {
			err = "ID cookie not found"
		}

		// get names for id
		names := namesStore.List(id)

		return c.Render("index", fiber.Map{
			"Error": err,
			"Names": names,
		})
	}
}

func PostGenerate(namesStore *store.NamesStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get id cookie
		id := c.Locals("id").(string)

		// input in form named "input" (idiomatic way to get form values in Fiber)
		input := strings.Clone(c.FormValue("input"))

		// validate input
		if input == "" {
			return c.Redirect("/")
		}

		// create good project name
		output := CreateProjectName(input)

		// generate names for id
		namesStore.Create(id, output)
		return c.Redirect("/")
	}
}
