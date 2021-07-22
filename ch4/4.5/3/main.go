package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Err      error
	Response *http.Response
}

// p.101 3つ以上のエラーが発生したら処理を停止する
func main() {
	checkStatus := func(done chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{
					Err:      err,
					Response: resp,
				}

				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()

		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{
		"https://www.google.com",
		"xxx",
		"https://badhost",
		"yyy",
		"https://example.com",
	}

	errCount := 0
	for result := range checkStatus(done, urls...) {
		if result.Err != nil {
			errCount++
			fmt.Printf("Err: %v\n", result.Err)

			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}

		fmt.Printf("Response: %v\n", result.Response.Status)
	}

}
