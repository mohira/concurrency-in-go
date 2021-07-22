package main

import (
	"fmt"
	"runtime"
	"time"
)

// p.93 ゴルーチンの親子間で親から子にキャンセルの信号を送れるようにする
// 親ゴルーチンがキャンセルしたいときに、キャンセルする
func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		// ゴルーチン1号
		go func() {
			// ゴルーチン1号が終了したのなら、このdeferが実行されるはず
			defer fmt.Println("（っ'-')╮ =͟͟͞(終了) ﾌﾞｫﾝ")
			defer close(terminated)

			for {
				select {

				case s := <-strings:
					fmt.Println(s)
				case <-done: // 親からキャンセルが来た場合は、returnする
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})

	// nilチャネルを渡しているけど、ちゃんとゴルーチンは終了する！
	terminated := doWork(done, nil)

	// ゴルーチン2号
	go func() {
		// 1秒後に操作をキャンセルする
		time.Sleep(1 * time.Second)
		fmt.Println("∠( ﾟдﾟ)／: Canceling doWork(ゴルーチン1号)")
		close(done)
	}()

	fmt.Printf("ゴルーチンの数: %d\n", runtime.NumGoroutine()) // 3

	<-terminated // ゴルーチン1号が終了されるまでブロック

	fmt.Printf("ゴルーチンの数: %d\n", runtime.NumGoroutine()) // 1

	fmt.Println("Mainゴルーチン おわり.")
}
