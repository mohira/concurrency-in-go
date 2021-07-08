package main

import (
	"fmt"
	"sync"
)

// p.40 2つのゴルーチンの間に合流ポイントを作った例
func main() {
	fmt.Println("mainゴルーチンすたーと")
	var wg sync.WaitGroup

	wg.Add(1)

	sayHello := func() {
		fmt.Println("hello")
		wg.Done()
	}

	go sayHello()
	wg.Wait()

	fmt.Println("mainゴルーチンおわり〜")

}
