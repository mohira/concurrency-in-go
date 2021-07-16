package main

import (
	"fmt"
	"time"
)

// defaultをfor-selectループで使う
// ゴルーチンの結果報告を待つ間に、他のゴルーチンで仕事を進められる！
func main() {

	done := make(chan interface{})
	go func() {
		fmt.Printf("（っ'-')╮ =͞  仕事するぞ！\n")
		time.Sleep(5 * time.Second)
		fmt.Printf("（っ'-')╮ =͞  仕事おわり！\n")
		close(done)
	}()

	fmt.Printf("∠( ﾟдﾟ)／ 他のゴルーチン（っ'-')╮ の仕事が終わるまで数を数えまくるぞい\n")
	var counter int
loop:
	for {

		select {
		case <-done:
			break loop
		default:
			// <-doneができない == doneチャネルが読み込めない(closeされていない)なら何もしない
			// つまり、Mainゴルーチンはcounterをインクリメントする仕事をする
			// 1つ仕事(インクリメント)しては、チェックする感じになる
		}
		counter++
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("∠( ﾟдﾟ)／ %d回数えたで！\n", counter)
}
