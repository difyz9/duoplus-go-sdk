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

	templates, err := client.Automation.ListOfficialTemplates(context.Background(), duoplus.TemplateListRequest{
		PaginationRequest: duoplus.PaginationRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range templates.List {
		fmt.Printf("%s %s\n", item.ID, item.Name)
	}
}