package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
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

	done := make(chan interface{})
	defer close(done)
	defer fmt.Printf("ゴルーチン数: %d\n", runtime.NumGoroutine())

	newResponses := func(done <-chan interface{}, urls []string) <-chan Result {
		results := make(chan Result)

		for _, url := range urls {
			go func(url string) {
				resp, err := http.Get(url)
				result := Result{
					Err:      err,
					Response: resp,
					Url:      url,
				}
				select {
				// selectを仕込んでちゃんとゴルーチンを終了できるようにしておく
				case <-done:
					return
				case results <- result:
				}
			}(url)
		}

		return results
	}

	results := newResponses(done, urls)

	for i := 0; i < len(urls); i++ {
		result := <-results

		if result.Err != nil {
			log.Println(result.Err)
			continue
		}

		fmt.Println(result.String())
	}
	fmt.Printf("ゴルーチン数: %d\n", runtime.NumGoroutine())

}
