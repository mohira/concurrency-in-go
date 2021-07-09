package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		// 結局 count という変数に対する操作権限を主張している
		// (deferがあるから挟んでいる感が最初は見えにくいけども)
		lock.Lock()
		defer lock.Unlock()
		count++

		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		// 結局 count という変数に対する操作権限を主張している
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// arithmetic は「算術」とか
	// əríθmətìk(米国英語)
	var arithmetic sync.WaitGroup

	// インクリメント
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)

		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// デクリメント
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)

		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete")
}
