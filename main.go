package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/yosssi/gohtml"
)

func main() {
	urlFlag := flag.String("u", "", "URL of the web page")
	flag.Parse()

	if *urlFlag == "" {
		flag.PrintDefaults()
		return
	}

	response, exception := http.Get(*urlFlag)
	handle(exception)

	sourceBinary, exception := io.ReadAll(response.Body)
	handle(exception)

	sourceString := string(sourceBinary)
	formattedSourceString := gohtml.Format(sourceString)

	fmt.Println(formattedSourceString)
}
