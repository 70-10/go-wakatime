package main

import (
	"encoding/json"
	"fmt"
	"time"

	wakatime "github.com/70-10/go-wakatime"
)

func main() {
	client := wakatime.NewClient("<YOUR API KEY>")
	today := time.Now().Format("2006-01-02")
	summaries, err := client.GetSummaries(today, today)
	if err != nil {
		panic(err)
	}

	jsonStr, err := prettyJSON(summaries)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonStr)
}

func prettyJSON(v interface{}) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
