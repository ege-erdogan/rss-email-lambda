package netutil

import (
	"io/ioutil"
	"log"
	"net/http"
)

// ReadFile returns the contents of the remote file as a string
func ReadFile(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	return string(content)
}
