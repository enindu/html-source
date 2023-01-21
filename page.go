package main

import (
	"os"
	"regexp"

	"github.com/yosssi/gohtml"
)

func downloadPage(name string, url string) []byte {
	// Download page
	filePath := download(name, url)

	// Read downloaded page
	fileBytes, exception := os.ReadFile(filePath)
	handle(exception)

	// Convert file bytes into string and remove HTML comments from it
	regex, exception := regexp.Compile(`<!--.*?-->`)
	handle(exception)

	fileString := regex.ReplaceAllString(string(fileBytes), "")

	// Format file string in prettier way and convert it into bytes
	sourceString := gohtml.Format(fileString)
	sourceBytes := []byte(sourceString)

	// Replace downloaded file with source bytes
	os.WriteFile(filePath, sourceBytes, 0644)

	// Return source bytes
	return sourceBytes
}
