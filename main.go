package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var projectName int64 = time.Now().Unix()
var projectPath string = fmt.Sprintf("downloads/%d", projectName)

func main() {
	// Create command line arguments
	urlFlag := flag.String("U", "", "URL of the web page\nEx: html-source -U <url>")

	flag.Parse()

	if *urlFlag == "" {
		flag.PrintDefaults()
		return
	}

	// Create project directory
	os.Mkdir(projectPath, 0755)

	// Download page
	sourceBytes := downloadPage("index.html", *urlFlag)

	// Download assets
	dowloadAssets(sourceBytes, *urlFlag)
}
