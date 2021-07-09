package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	// 絶対終了しないゴルーチンくん
	noop := func() {
		// チャンネル受信をずーっと待っている切ないゴルーチン
		// noop は 「何もしない命令」という意味。"No Operation"らしい
		wg.Done()
		<-c
	}

	// 生成するゴルーチンの数
	const numGoroutines = 1e5

	wg.Add(numGoroutines)

	before := memConsumed()

	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()

	after := memConsumed()

	fmt.Printf("%.3fKB\n", float64(after-before)/numGoroutines/1000)

}
