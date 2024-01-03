package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	key := os.Getenv("MASTODON_KEY")

	form := url.Values{}
	form.Add("status", "Hello World!")
	form.Add("visibility", "public")
	form.Add("language", "eng")

	request, err := http.NewRequest("POST", "https://mastodon.mallegolhansen.com/api/v1/statuses", strings.NewReader(form.Encode()))

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
