package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpClient struct {
	RequestUrl string
}

func NewHttpClient(requestUrl string) *HttpClient {
	return &HttpClient{
		RequestUrl: requestUrl,
	}
}

func (h HttpClient) SendRequest() (*CategoryTreeResponse, error) {
	resp, err := http.Get(h.RequestUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return nil, err
	}

	var categoryTree CategoryTreeResponse
	err = json.NewDecoder(resp.Body).Decode(&categoryTree)
	if err != nil {
		return &categoryTree, err
	}

	return &categoryTree, err
}
