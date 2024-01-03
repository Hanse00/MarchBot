package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	key := os.Getenv("MASTODON_KEY")
	masto := mastodon{"https://mastodon.mallegolhansen.com", key}

	post := post{"*tap tap* Hi?", "public", "eng"}

	response, err := masto.post(post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
