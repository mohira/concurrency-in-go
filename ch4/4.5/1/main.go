package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)

		go func() {
			defer close(responses)

			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					// ① このゴルーチンができるのはエラーを出力するところまで！
					// 「親Goroutineに、エラーを返して、親が処理する」みたいなことができない！
					// エラーを返せないので、**誰かが処理してくれることを願うことしかできない**
					fmt.Println(err)
					continue
				}

				select {
				case <-done:
					return
				case responses <- resp:
				}
			}

		}()

		return responses
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

	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
