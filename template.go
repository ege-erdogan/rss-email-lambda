package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"time"
)

const templatesPath = "template/"

// GenerateMessage creates main message
func GenerateMessage(blocks []string) string {
	var htmlBlocks []template.HTML
	htmlBlocks = append(htmlBlocks, template.HTML(GenerateHeader()))
	for _, text := range blocks {
		htmlBlocks = append(htmlBlocks, template.HTML(text))
	}
	return executeTemplate(templatesPath+"main.html", htmlBlocks)
}

// GenerateHeader creates header HTML for feed
func GenerateHeader() string {
	date := time.Now().Format("January 2, 2006")
	return executeTemplate(templatesPath+"header.html", date)
}

// GenerateHTMLFeedBlock create HTML template for one feed
func GenerateHTMLFeedBlock(feed Feed) string {
	return executeTemplate(templatesPath+"feed.html", feed)
}

func executeTemplate(filePath string, data interface{}) string {
	content, _ := ioutil.ReadFile(filePath)
	tmpl, err := template.New(filePath).Parse(string(content))
	check(err)

	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)
	check(err)
	return doc.String()
}

// GetCurrentDate returns current date as formatted string
func GetCurrentDate() string {
	return time.Now().Format("January 2, 2020")
}
