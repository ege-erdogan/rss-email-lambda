package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

const days = 7
const feedsURL = "https://raw.githubusercontent.com/ege-erdogan/rss-email/master/feeds.txt"

func main() {
	dateThreshold := time.Now().AddDate(0, 0, -days)
	msg := "<html><h1>RSS FEEDS</h1> \n"

	resp, err := http.Get(feedsURL)
	check(err)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	check(err)

	urls := strings.Split(string(data), "\n")

	htmlChannel := make(chan string)

	for i, url := range urls {
		fmt.Printf("[%d] [%s]\n", i, url)
	}

	for _, url := range urls {
		go func(url string) {
			htmlChannel <- fetch(url, dateThreshold)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// FIXME: blocks if a fetch call errs
		msg += <-htmlChannel
	}

	msg += "</html>\n\n"
	send(os.Getenv("RSS_TARGET"), msg)
}

func fetch(url string, threshold time.Time) string {
	posts := make(map[string]string)

	fmt.Println(url)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	check(err)

	for i := 0; i < len(feed.Items); i++ {
		if feed.Items[i].PublishedParsed.After(threshold) {
			title := feed.Items[i].Title
			link := feed.Items[i].Link
			posts[title] = link
		}
	}

	return GenerateHTMLFeedBlock(feed.Title, posts)
}

func send(to, body string) {
	from := os.Getenv("GMAIL_NAME")
	password := os.Getenv("GMAIL_PASS")

	msg := "From: " + from + "\n"
	msg += "To: " + to + "\n"
	msg += "Content-Type: text/html\n"
	msg += "Subject: RSS FEEDS\n\n"
	msg += body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
