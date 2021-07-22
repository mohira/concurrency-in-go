package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan int)

		// selectを使っているので、ゴルーチンはちゃんと終了する
		go func() {
			select {
			case ch <- 123:
			default:
			}
		}()

		go func() {
			select {
			case ch <- 123:
			default:
			}
		}()

		// 少しsleepして、2つのゴルーチンがdefaultで終わるようにする
		time.Sleep(100 * time.Millisecond)

		<-ch

		fmt.Fprintf(w, "NumGoroutine: %d\n", runtime.NumGoroutine())
	})

	http.ListenAndServe(":8080", nil)
}
