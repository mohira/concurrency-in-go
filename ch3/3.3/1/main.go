package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		// 無名ゴルーチンは、キャパシティが空くまで待機する
		// キャパシティがあれば、チャネルに書き込む
		ch <- "Hello channels!"
	}()

	// メインゴルーチンは、読み込みをする
	// もし、チャネルが空だったら、メインゴルーチンは待機する
	// つまり、チャネルにデータが書き込まれるまで、メインゴルーチンは待機する
	// だから、プログラムは終了しないし、実行結果が「決定的」になる
	fmt.Println(<-ch)
}
