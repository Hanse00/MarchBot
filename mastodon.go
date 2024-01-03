package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type mastodon struct {
	baseUrl string
	apiKey  string
}

type post struct {
	message    string
	visibility string
	language   string
}

func (m mastodon) post(p post) (string, error) {
	post_url, err := url.JoinPath(m.baseUrl, "api/v1/statuses")
	fmt.Println(post_url)
	if err != nil {
		return "", err
	}

	form := url.Values{}
	form.Add("status", p.message)
	form.Add("visibility", p.visibility)
	form.Add("language", p.language)

	request, err := http.NewRequest("POST", post_url, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.apiKey))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}
	return string(body), nil
}
