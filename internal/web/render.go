package web

import (
	"log"
	"strings"
)

func CreateProjectName(input string) string {
	// lowercase and replace spaces with _
	output := strings.ToLower(input)
	output = strings.ReplaceAll(output, " ", "_")
	log.Printf("Generated project name: %s from input: %s", output, input)
	return output
}
