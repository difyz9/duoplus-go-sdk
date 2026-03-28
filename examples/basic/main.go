package main

import (
	"context"
	"fmt"
	"log"
	"os"

	duoplus "duoplus-go-sdk"
)

func main() {
	apiKey := os.Getenv("DUOPLUS_API_KEY")
	if apiKey == "" {
		log.Fatal("DUOPLUS_API_KEY is required")
	}

	client, err := duoplus.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.CloudPhones.List(context.Background(), duoplus.CloudPhoneListRequest{
		PaginationRequest: duoplus.PaginationRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range resp.List {
		fmt.Printf("%s %s %d\n", item.ID, item.Name, item.Status)
	}
}