package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

const templatesPath = "template/"

// GenerateMessage creates main message
func GenerateMessage(blocks []string) string {
	var htmlBlocks []template.HTML
	for _, text := range blocks {
		htmlBlocks = append(htmlBlocks, template.HTML(text))
	}
	return executeTemplate(templatesPath+"main.html", htmlBlocks)
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
