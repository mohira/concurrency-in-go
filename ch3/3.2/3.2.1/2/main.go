package main

import (
	"fmt"
	"sync"
	"time"
)

// あえて競合状態をやってみる
func main() {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1) // Addの呼び出しをゴルーチン内部でやってしまっている！ → 競合状態！
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1 * time.Second)
	}()

	go func() {
		wg.Add(1) // Addの呼び出しをゴルーチン内部でやってしまっている！ → 競合状態！
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All goroutine complete")
}
