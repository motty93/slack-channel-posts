package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

func main() {
	token := ""
	channelID := ""
	dateStr := "2006-01-02 15:04:05"

	oldest, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Fatalf("Failed to parse date: %v", err)
	}

	api := slack.New(token)
	params := &slack.GetConversationHistoryParameters{
		ChannelID: channelID,
		Oldest:    fmt.Sprintf("%f", float64(oldest.Unix())),
		Limit:     100,
	}

	res, err := api.GetConversationHistory(params)
	if err != nil {
		log.Fatalf("Failed to get conversation history: %s", err)
	}

	for _, message := range res.Messages {
		if message.Text != "" && message.User != "" {
			ts, err := strconv.ParseFloat(message.Timestamp, 64)
			if err != nil {
				log.Fatalf("Failed to parse timestamp: %v", err)
				continue
			}
			postDate := time.Unix(int64(ts), 0).Format("2006-01-02 15:04:05")
			fmt.Printf("Date: %s, User: %s, Text: %s\n", postDate, message.User, message.Text)
		}
	}
}
