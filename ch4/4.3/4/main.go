package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// p.95
func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream) // ところで、closeするのは所有者の責任だよね

			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	fmt.Printf("ゴルーチンの数: %d\n", runtime.NumGoroutine())

	// doneチャネルの所有者はMainゴルーチン
	// なので、責任を持ってcloseする
	close(done)

	// ゴルーチンのdeferの処理を待つためのSleep
	time.Sleep(1 * time.Second)

	fmt.Printf("ゴルーチンの数: %d\n", runtime.NumGoroutine())
}
