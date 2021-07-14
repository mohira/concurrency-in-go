package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	begin := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			fmt.Printf("\t（っ'-')╮ =͞ ゴルーチン%d号 待機...\n", n)

			// 狼煙がくるまで、各ゴルーチンは待機
			// beginチャネルがcloseされるまで(受信可能になるまで)待機
			// close(begin)されたあとなら、無限に受信ができる(そして、ゼロ値を返す)
			<-begin

			fmt.Printf("\t（っ'-')╮ =͞ ゴルーチン%d号 出動！\n", n)
		}(i)
	}

	// ゴルーチン間のコミュニケーションをイメージしやすくするための仕込み
	fmt.Println("∠( ﾟдﾟ)／: 狼煙の準備なう...")
	time.Sleep(1 * time.Second)
	fmt.Println("∠( ﾟдﾟ)／: 狼煙だ〜〜〜")
	time.Sleep(1 * time.Second)

	close(begin) // 終了の狼煙を上げる

	wg.Wait() // 無名ゴルーチンたちが終了するまで、Mainゴルーチンは待機

	fmt.Println("∠( ﾟдﾟ)／: おしまい")
}
