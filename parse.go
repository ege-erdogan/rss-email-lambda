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

	msg := "RSS FEEDS \n"
	lines := 0

	start := time.Now()

	receive := make(chan string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		lines++
		go fetch(url, dateThreshold, receive)
	}

	for i := 0; i < lines; i++ {
		msg += <-receive
	}

	fmt.Println(time.Since(start).Seconds())
	fmt.Println(msg)
	// send("ege@erdogan.dev", msg)
}

func fetch(url string, threshold time.Time, out chan string) {
	msg := ""

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	check(err)

	msg += "\n" + feed.Title + "\n"
	for i := 0; i < len(feed.Items); i++ {
		if feed.Items[i].PublishedParsed.After(threshold) {
			msg += "\t" + feed.Items[i].Published + "  -  " + feed.Items[i].Title + "\n"
		}
	}

	out <- msg
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
