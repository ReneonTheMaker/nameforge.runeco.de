package web

import (
	"strconv"
	"strings"

	"app/internal/model"
	"app/internal/store"

	"github.com/gofiber/fiber/v2"
)

// index page handler
func GetMain(namesStore *store.NamesStore, configStore *store.RenderConfigStore) fiber.Handler {
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

		// get config for id
		config, _ := configStore.Get(id)

		return c.Render("index", fiber.Map{
			"Error":  err,
			"Names":  names,
			"Config": config,
		})
	}
}

func PostGenerate(namesStore *store.NamesStore, configStore *store.RenderConfigStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get id cookie
		id := c.Locals("id").(string)

		// input in form named "input" (idiomatic way to get form values in Fiber)
		input := strings.Clone(c.FormValue("input"))

		// form config vals
		isFileName := c.FormValue("file_format") == "on"
		isLowerCase := c.FormValue("lower_case") == "on"
		isProjectName := c.FormValue("project") == "on"
		versionNumber, err := strconv.Atoi(c.FormValue("version"))
		extension := c.FormValue("extension")
		if err != nil {
			versionNumber = 0 // default to 0 if parsing fails
		}

		config := model.RenderConfig{
			FileName:      isFileName,
			Lowercase:     isLowerCase,
			Project:       isProjectName,
			VersionNumber: versionNumber,
			FileExtension: extension,
		}

		// validate input
		if input == "" {
			return c.Redirect("/")
		}

		// create good project name
		output := CreateProjectName(input, config)

		// generate names for id
		namesStore.Create(id, output)

		// save config for id
		configStore.Set(id, config)

		return c.Redirect("/")
	}
}
