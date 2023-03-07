package api

import (
	"net/http"

	"github.com/muhammad-asn/idcloudhost-go-lib/s3"
)

type APIClient struct {
	S3 *s3.S3Api
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewClient(authToken string) (*APIClient, error) {
	c := http.Client{}
	var ApiClient = APIClient{
		S3: &s3.S3Api{},
	}

	ApiClient.S3.Init(&c, authToken)
	return &ApiClient, nil
}
