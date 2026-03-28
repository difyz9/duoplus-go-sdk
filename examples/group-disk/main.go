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

	groups, err := client.Groups.List(context.Background(), duoplus.GroupListRequest{Page: 1})
	if err != nil {
		log.Fatal(err)
	}

	files, err := client.CloudDisk.List(context.Background(), duoplus.CloudDiskListRequest{
		PaginationRequest: duoplus.PaginationRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("groups=%d files=%d\n", len(groups.List), len(files.List))
}