package main

import (
	"fmt"
)

// p.93 簡単なゴルーチンリークの例
func main() {
	doWork := func(strings chan string) <-chan interface{} {
		completed := make(chan interface{})

		// ゴルーチン1号: （っ'-')╮
		go func() {
			// ゴルーチン1号が終了したのなら、このdeferが実行されるはず
			defer fmt.Println("（っ'-')╮ =͟͟͞(終了) ﾌﾞｫﾝ")
			defer close(completed)

			for s := range strings {
				fmt.Println(s)
			}
		}()

		return completed
	}

	// Mainゴルーチンでは、nilチャネル を渡しているので、stringsには絶対書き込みは起きない
	// => ゴルーチン1号はメモリ内の残り続ける(このプロセスが生きている限り)
	doWork(nil)

	fmt.Println("Mainゴルーチン おわり.")
}
