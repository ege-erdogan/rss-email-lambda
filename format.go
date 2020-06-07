package main

import "fmt"

// GenerateHTMLFeedBlock create HTML template for one feed
func GenerateHTMLFeedBlock(feedTitle string, posts map[string]string) string {
	html := fmt.Sprintf("<h3>%s</h3><ul>", feedTitle)
	for title, link := range posts {
		html += "<li>"
		html += fmt.Sprintf("<a href=\"%s\">%s</a></li>\n", link, title)
	}
	html += "</ul>"

	return html
}
