package main

import (
	"context"
	"fmt"
	"log"
	"os"

	duoplus "duoplus-go-sdk"
	"duoplus-go-sdk/clouddisk"
	"duoplus-go-sdk/common"
	"duoplus-go-sdk/group"
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

	groups, err := client.Groups.List(context.Background(), group.ListRequest{Page: 1})
	if err != nil {
		log.Fatal(err)
	}

	files, err := client.CloudDisk.List(context.Background(), clouddisk.ListRequest{
		PaginationRequest: common.PaginationRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("groups=%d files=%d\n", len(groups.List), len(files.List))
}
