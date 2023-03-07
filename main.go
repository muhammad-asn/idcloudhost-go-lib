package main

import (
	"fmt"
	"log"
	"os"

	"github.com/muhammad-asn/idcloudhost-go-lib/api"
	"github.com/muhammad-asn/idcloudhost-go-lib/s3"
)

func main() {
	authToken := os.Getenv("AUTH_TOKEN")

	if authToken == "" {
		fmt.Println("Environment variable AUTH_TOKEN is not set.")
		return
	}

	s3bucket := s3.S3Bucket{
		Name:             "test-bucket-terraform-3",
		BillingAccountId: 1200190928,
	}

	client, err := api.NewClient(authToken)

	if err != nil {
		log.Fatal(err)
	}

	// Create
	client.S3.Create(s3bucket)

	// Get
	client.S3.Get(s3bucket)

	// Delete
	client.S3.Delete(s3bucket)
}
