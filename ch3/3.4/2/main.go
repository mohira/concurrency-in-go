package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	c := make(chan interface{})

	go func() {
		fmt.Printf("（っ'-')╮ =͞  5秒寝まーす\n")
		time.Sleep(5 * time.Second)
		fmt.Printf("（っ'-')╮ =͞  close!!!\n")
		close(c)
	}()

	fmt.Println("∠( ﾟдﾟ)／ 読み込みを待機...")

	// ホントは <-c だけで十分(selectなくてもいける)けど、記述を拡張していくためにわざと書いている
	select {
	// チャネルに読み込み可能になる(書き込まれるか、closeされるか)まで、チャネルはブロックする
	// つまり、Mainゴルーチンが待機するってこと
	case <-c:
		fmt.Printf("∠( ﾟдﾟ)／ %v経ちましたとさ\n", time.Since(start))
	}
}
