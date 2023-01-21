package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func dowloadAssets(sourceBytes []byte, url string) {
	// Get sanitized links
	sanitizedLinks := getSanitizedLinks(sourceBytes, url)

	// Download assets
	for link, sanitizedLink := range sanitizedLinks {
		linkIndex := strings.LastIndex(link, "/")
		if linkIndex != -1 {
			fileName := link[linkIndex+1:]
			filePath := link[:linkIndex+1]

			_, exception := os.Stat(filePath)
			if os.IsNotExist(exception) {
				os.MkdirAll(projectPath+"/"+filePath, 0755)
			}

			download(filePath+fileName, sanitizedLink)
			continue
		}

		download(link, sanitizedLink)
	}
}

func getSanitizedLinks(sourceBytes []byte, url string) map[string]string {
	// Get links
	links := getLinks(sourceBytes)

	// Sanitize links
	regex, exception := regexp.Compile(`^(https?|ftp|file)://|//`)
	handle(exception)

	sanitizedLinks := make(map[string]string)

	for _, link := range links {
		if regex.MatchString(link) {
			sanitizedLinks[link] = link
			continue
		}

		sanitizedLinks[link] = fmt.Sprintf("%s/%s", url, link)
	}

	// Return sanitized links
	return sanitizedLinks
}

func getLinks(sourceBytes []byte) []string {
	var links []string

	// Convert source bytes into reader
	sourceReader := bytes.NewReader(sourceBytes)

	// Convert source reader into HTML document and get links from it
	htmlDocument, exception := goquery.NewDocumentFromReader(sourceReader)
	handle(exception)

	htmlDocument.Find("link, img, script, audio, video").Each(func(_ int, selection *goquery.Selection) {
		// Get href links
		hrefLink, hrefLinkExists := selection.Attr("href")
		if hrefLinkExists {
			links = append(links, hrefLink)
		}

		// Get src links
		srcLink, srcLinkExists := selection.Attr("src")
		if srcLinkExists {
			links = append(links, srcLink)
		}
	})

	// Return links
	return links
}
