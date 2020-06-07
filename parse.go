package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	file, err := os.Open("feeds.txt")
	check(err)
	defer file.Close()

	fp := gofeed.NewParser()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		feed, _ := fp.ParseURL(scanner.Text())
		fmt.Println(feed.Title)
		for i := 0; i < len(feed.Items); i++ {
			fmt.Printf("\t%s - %s\n", feed.Items[i].Published, feed.Items[i].Title)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
