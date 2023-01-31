package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type HttpClient struct {
	RequestUrl string
}

func NewHttpClient(requestUrl string) *HttpClient {
	return &HttpClient{
		RequestUrl: requestUrl,
	}
}

func (h HttpClient) SendRequest() {
	// start := time.Now()
	resp, err := http.Get(h.RequestUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	var categoryTree CategoryTreeResponse
	err = json.NewDecoder(resp.Body).Decode(&categoryTree)
	if err != nil {
		panic(err)
	}

	// fmt.Println(len(categoryTree.Categories))
	// recursivePrint(categoryTree.Categories)

	// fmt.Printf("took: %d ns\n", time.Since(start).Nanoseconds())
}

func recursivePrint(r []*Category) {
	for _, c := range r {
		fmt.Println(c.Name)
		recursivePrint(c.SubCategories)
	}
}
