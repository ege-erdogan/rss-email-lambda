package main

// Feed encapsulates basic data about an RSS feed
// 	to be passed onto the template engine
type Feed struct {
	Title string
	Link  string
	Posts []Post
}

// Post encapsulates basic data about a blog post
type Post struct {
	Title      string
	Link       string
	DateString string
}
