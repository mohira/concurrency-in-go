package main

import (
	"fmt"
	"net/http"
	"runtime"
)

// $ for i in {1..10}; do curl "localhost:8080"; done
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan int)

		go func() {
			ch <- 456
		}()
		go func() {
			ch <- 123
		}()

		<-ch

		fmt.Fprintf(w, "NumGoroutine: %d\n", runtime.NumGoroutine())
	})

	http.ListenAndServe(":8080", nil)
}
