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
		Name:             "test-bucket-terraform-2",
		BillingAccountId: 1200190928,
	}

	client, err := api.NewClient(authToken)

	if err != nil {
		log.Fatal(err)
	}

	// Create
	fmt.Println(client.S3.Create(s3bucket))

	// Delete
	fmt.Println(client.S3.Delete(s3bucket))
}
