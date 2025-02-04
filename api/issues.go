package api

import (
	"encoding/json"
	"log"
)

// TODO: don't hardcode language? Take from config
const URL = "https://api.irail.be/v1/disturbances/?format=json&lineBreakCharacter=''&lang=nl"

type DescriptionLink struct {
	ID   string `json:"id"`
	Link string `json:"link"`
	Text string `json:"text"`
}

type DescriptionLinks struct {
	Links []DescriptionLink `json:"DescriptionLink"`
}

type Issue struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Type        string           `json:"type"`
	Timestamp   string           `json:"timestamp"`
	Links       DescriptionLinks `json:"descriptionLinks"`
}

func GetIssues() []Issue {

	body, _ := makeAPIRequest(URL)

	return ParseIssues(body)

}

func ParseIssues(jsonData []byte) []Issue {
	var result struct {
		Stations []Issue `json:"disturbance"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %v - input data: %s", err, string(jsonData))
	}
	return result.Stations
}
