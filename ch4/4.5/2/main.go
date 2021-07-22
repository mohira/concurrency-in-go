package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				resp, err := http.Get(url)

				// if err != nil のチェックは、ゴルーチン内でやらない！
				result := Result{
					Error:    err,
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

	for result := range checkStatus(done, urls...) {
		// ゴルーチンの中(相対的に"小さな"コンテキスト)で起きたエラーを
		// "大きな"コンテキストで扱えている！
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
			continue
		}

		fmt.Printf("Response: %v\n", result.Response.Status)
	}

}
