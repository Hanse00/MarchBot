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

type mastodon struct {
	baseUrl string
	apiKey  string
}

func (m mastodon) post() string {
	post_url, err := url.JoinPath(m.baseUrl, "api/v1/statuses")
	fmt.Println(post_url)
	if err != nil {
		log.Fatal(err)
	}

	form := url.Values{}
	form.Add("status", "Hello World!")
	form.Add("visibility", "public")
	form.Add("language", "eng")

	request, err := http.NewRequest("POST", post_url, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.apiKey))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func main() {
	key := os.Getenv("MASTODON_KEY")
	masto := mastodon{"https://mastodon.mallegolhansen.com", key}

	response := masto.post()
	fmt.Println(response)
}
