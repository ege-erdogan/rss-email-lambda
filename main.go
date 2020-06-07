package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

const days = 7

func main() {
	file, err := os.Open("feeds.txt")
	check(err)
	defer file.Close()

	dateThreshold := time.Now().AddDate(0, 0, -days)

	msg := "<html><h1>RSS FEEDS</h1> \n"
	lines := 0

	start := time.Now()

	htmlChannel := make(chan string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		lines++
		go func(url string) {
			htmlChannel <- fetch(url, dateThreshold)
		}(url)
	}

	for i := 0; i < lines; i++ {
		msg += <-htmlChannel
	}

	msg += "</html>\n\n"
	fmt.Println(time.Since(start).Seconds())
	fmt.Println(msg)
	send("ege@erdogan.dev", msg)
}

func fetch(url string, threshold time.Time) string {
	posts := make(map[string]string)

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
