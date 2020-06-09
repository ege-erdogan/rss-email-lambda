package main

import (
	"bytes"
	"html/template"
	"os"
	"time"

	"./netutil"
)

var templatesPath = os.Getenv("TEMPLATE_S3_BUCKET")

// GenerateMessage creates main message
func GenerateMessage(blocks []string) string {
	var htmlBlocks []template.HTML
	htmlBlocks = append(htmlBlocks, template.HTML(GenerateHeader()))
	for _, text := range blocks {
		htmlBlocks = append(htmlBlocks, template.HTML(text))
	}
	return executeTemplate("main.html", htmlBlocks)
}

// GenerateHeader creates header HTML for feed
func GenerateHeader() string {
	date := time.Now().Format("January 2, 2006")
	return executeTemplate("header.html", date)
}

// GenerateHTMLFeedBlock create HTML template for one feed
func GenerateHTMLFeedBlock(feed Feed) string {
	return executeTemplate("feed.html", feed)
}

func executeTemplate(templateName string, data interface{}) string {
	templateHTML := netutil.ReadFile(templatesPath + templateName)
	tmpl, err := template.New(templateName).Parse(string(templateHTML))
	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)

	return doc.String()
}

// GetCurrentDate returns current date as formatted string
func GetCurrentDate() string {
	return time.Now().Format("January 2, 2020")
}
