package main

import (
	"fmt"
	"sync"
)

// https://pkg.go.dev/sync#Once
func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	// 単一チャネルで
	// 『Go言語による並行処理』では sync.WaitGroup 使ってた
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
