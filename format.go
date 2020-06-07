package main

import "fmt"

// GenerateHTMLFeedBlock create HTML template for one feed
func GenerateHTMLFeedBlock(feedTitle string, posts map[string]string) string {
	html := fmt.Sprintf("<h3>%s</h3><ul>", feedTitle)
	for title, link := range posts {
		html += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", link, title)
	}
	html += "</ul>"

	return html
}
