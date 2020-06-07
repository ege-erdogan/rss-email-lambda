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

	fp := gofeed.NewParser()
	msg := "RSS FEEDS \n"

	start := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		feed, _ := fp.ParseURL(scanner.Text())
		msg += "\n" + feed.Title + "\n"
		for i := 0; i < len(feed.Items); i++ {
			if feed.Items[i].PublishedParsed.After(dateThreshold) {
				msg += "\t" + feed.Items[i].Published + "  -  " + feed.Items[i].Title + "\n"
			}
		}
	}

	fmt.Println(time.Since(start).Seconds())
	// send("ege@erdogan.dev", msg)
}

func send(to, body string) {
	from := os.Getenv("GMAIL_NAME")
	pass := os.Getenv("GMAIL_PASS")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: RSS Feeds\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	check(err)

	fmt.Printf("Sent mail to %s\n", to)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
