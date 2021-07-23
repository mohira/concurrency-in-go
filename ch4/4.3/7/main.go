package main

import (
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Err      error
	Response *http.Response
	Url      string
}

func (r *Result) String() string {
	return fmt.Sprintf("%s: %s", r.Url, r.Response.Status)
}

func main() {
	urls := []string{
		"https://example.com",
		"https://python.org",
		"https://google.com",
		"https://badhost",
	}

	newResults := func(urls []string) <-chan Result {
		results := make(chan Result)

		for _, url := range urls {
			go func(url string) {
				resp, err := http.Get(url)
				result := Result{
					Err:      err,
					Response: resp,
					Url:      url,
				}

				// 素直に送信しているが...
				results <- result

			}(url)
		}

		return results
	}

	results := newResults(urls)

	// 本当に終了条件のないfor文になってしまっている
	for result := range results {
		if result.Err != nil {
			log.Println(result.Err)
			continue
		}

		fmt.Println(result.String())
	}

}
