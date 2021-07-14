package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		// チャネルへの書き込みが起こらないように仕掛ける
		if true {
			return
		}
		ch <- "Hello channels!"
	}()

	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
}
