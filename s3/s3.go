package s3

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type S3Api struct {
	c              HTTPClient
	AuthToken      string
	BillingAccount int
	ApiEndpoint    string
	S3Bucket       S3Bucket
}
type S3Bucket struct {
	Name             string    `json:"name"`
	SizeBytes        int       `json:"size_bytes"`
	BillingAccountId int       `json:"billing_account_id"`
	NumObjects       int       `json:"num_objects"`
	CreatedAt        time.Time `json:"created_at"`
	ModifiedAt       time.Time `json:"modified_at"`
	IsSuspended      bool      `json:"is_suspended"`
}

func (s3 *S3Api) Init(c HTTPClient, authToken string) error {

	s3.c = c
	s3.AuthToken = authToken
	s3.ApiEndpoint = fmt.Sprintf("https://api.idcloudhost.com/v1/storage/bucket")

	r, err := http.Get(s3.ApiEndpoint)

	if err != nil {
		log.Fatal(err)
	}

	if r.StatusCode == http.StatusNotFound {
		return fmt.Errorf("S3 Not Found")
	}

	return nil
}

func (s3 *S3Api) Create(sb S3Bucket) error {

	data := url.Values{}

	data.Set("name", sb.Name)
	data.Set("billing_account_id", strconv.Itoa(sb.BillingAccountId))

	req, err := http.NewRequest("PUT", s3.ApiEndpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apikey", s3.AuthToken)

	r, err := s3.c.Do(req)

	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("%v", r.StatusCode))
	}
	return json.NewDecoder(r.Body).Decode(&s3.S3Bucket)

}

func (s3 *S3Api) Modify(sb S3Bucket) error {

	data := url.Values{}

	data.Set("name", sb.Name)
	data.Set("billing_account_id", strconv.Itoa(sb.BillingAccountId))

	req, err := http.NewRequest("PATCH", s3.ApiEndpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apikey", s3.AuthToken)

	r, err := s3.c.Do(req)

	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("%v", r.StatusCode))
	}
	return json.NewDecoder(r.Body).Decode(&s3.S3Bucket)

}

func (s3 *S3Api) Delete(sb S3Bucket) error {
	data := url.Values{}

	data.Set("name", sb.Name)

	req, err := http.NewRequest("DELETE", s3.ApiEndpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apikey", s3.AuthToken)

	r, err := s3.c.Do(req)

	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("%v", r.StatusCode))
	}
	return json.NewDecoder(r.Body).Decode(&s3.S3Bucket)

}
