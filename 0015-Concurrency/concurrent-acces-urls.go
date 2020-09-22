package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://golang.org/",
	"https://github.com/",
	"https://abcde123.com/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func main() {
	resp := httpGet(urls)
	result(resp)
}

func httpGet(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	client := http.Client{}

	responses := []*HttpResponse{}

	for _, url := range urls {
		go func(url string) {

			fmt.Println("Fetching url", url)
			resp, err := client.Get(url)

			ch <- &HttpResponse{
				url:      url,
				response: resp,
				err:      err,
			}

		}(url)
	}

		for {
			select {
			case r := <-ch:
				fmt.Printf("%s was fetched\n", r.url)
				if r.err != nil {
					fmt.Println("with an error", r.err)
				}
				responses = append(responses, r)
				if len(responses) == len(urls) {
					return responses
				}

			case <-time.After(10 * time.Millisecond):
				fmt.Printf(".")
			}
		}


	return responses

}

func result(resp []*HttpResponse) {
	for _, result := range resp {
		if result != nil && result.response != nil {
			fmt.Println(result.response.Status)
		} else {
			fmt.Println(result.url, "not found")
		}
	}

}
