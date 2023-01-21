package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download(name string, url string) string {
	// Create HTTP request to the URL and get HTTP response
	response, exception := http.Get(url)
	handle(exception)

	defer response.Body.Close()

	// Convert HTTP response into bytes
	sourceBytes, exception := io.ReadAll(response.Body)
	handle(exception)

	// Create directories and files using source bytes
	filePath := fmt.Sprintf(projectPath+"/%s", name)

	os.WriteFile(filePath, sourceBytes, 0644)

	// Return file path
	return filePath
}
