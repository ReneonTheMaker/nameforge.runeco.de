package web

import (
	"strconv"
	"strings"

	"app/internal/model"
)

func CreateProjectName(input string, config model.RenderConfig) string {
	if config.VersionNumber > 0 {
		input = input + " v" + strconv.Itoa(config.VersionNumber)
	}
	if config.Project {
		input = "Project " + input
	}
	if config.Lowercase {
		input = strings.ToLower(input)
	}
	if config.FileName {
		input = strings.ReplaceAll(input, " ", "_")
	}
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "'", "")
	config.FileExtension = strings.TrimSpace(config.FileExtension)
	if config.FileExtension != "" {
		input = input + "." + config.FileExtension
	}
	return input
}
