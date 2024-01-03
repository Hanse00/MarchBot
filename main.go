package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func formatNum(num int) string {
	last := num % 10
	postfix := ""
	switch last {
	case 1:
		postfix = "st"
	case 2:
		postfix = "nd"
	case 3:
		postfix = "rd"
	default:
		postfix = "th"
	}

	last_two := num % 100
	if last_two >= 4 && last_two <= 20 {
		postfix = "th"
	}

	return fmt.Sprintf("%d%s", num, postfix)
}

func marchDate() int {
	then := time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local)
	diff := time.Since(then)
	return int(diff.Hours()/24) + 1
}

func fmtMessage(date int) string {
	return fmt.Sprintf("Today's date is March %s 2020.", formatNum(date))
}

func main() {
	key := os.Getenv("MASTODON_KEY")
	url := os.Getenv("MASTODON_URL")
	masto := mastodon{url, key}

	post := post{fmtMessage(marchDate()), "public", "eng"}

	_, err := masto.post(post)
	if err != nil {
		log.Fatal(err)
	}
}
