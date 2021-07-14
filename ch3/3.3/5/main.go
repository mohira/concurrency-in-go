package main

import (
	"fmt"
	"strings"
	"time"
)

// ゴルーチン間のコミュニケーション感を演出してみる
// Mainゴルーチン と 無名ゴルーチン にアイコンを付ける → ∠( ﾟдﾟ)／  っ'-')╮
// 発言させてみる。sleepをつけると会話っぽくなる
// https://zenn.dev/mohira/scraps/95427e91e27c93#comment-f8da2b7318c040
func main() {
	ch := make(chan int)

	go func() {
		defer close(ch) // 無名ゴルーチンでforループの処理が終わったら、チャネルを閉じる

		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("無名ゴルーチン（っ'-')╮ =͟͟͞(%d) : データ書き込んだよ！\n", i)
		}
	}()

	for integer := range ch {
		time.Sleep(1 * time.Second)

		fmt.Printf("Mainゴルーチン ∠( ﾟдﾟ)／: (%d) を受け取ったぞ\n", integer)
		time.Sleep(1 * time.Second)

		fmt.Printf("Mainゴルーチン ∠( ﾟдﾟ)／: 次、送って良いぞ〜\n")
		time.Sleep(1 * time.Second)
		fmt.Println(strings.Repeat("-", 30))
	}

	fmt.Println("おしまい")
}
