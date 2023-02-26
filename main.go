package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/muhammad-asn/idcloudhost-go-lib/s3"
)

func main() {
	c := http.Client{}
	authToken := os.Getenv("AUTH_TOKEN")

	if authToken == "" {
		fmt.Println("Environment variable AUTH_TOKEN is not set.")
		return
	}

	s3bucket := s3.S3Bucket{
		Name:             "test-bucket-terraform-2",
		BillingAccountId: 1200190928,
	}

	s3api := s3.S3Api{}
	s3api.Init(&c, authToken)

	// Create
	s3api.Create(s3bucket)

	// Delete
	s3api.Delete(s3bucket)
}
