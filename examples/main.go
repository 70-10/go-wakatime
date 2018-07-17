package main

import (
	"encoding/json"
	"fmt"

	wakatime "github.com/70-10/go-wakatime"
)

func main() {
	client := wakatime.NewClient("<YOUR API KEY>")
	summaries, err := client.GetLast7Days()
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
