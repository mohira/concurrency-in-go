package main

import "fmt"

func main() {
	// chanOwner は intChチャネル の所有権を持っている
	// で、「非所有者」が使えるように、「読み込み専用チャネル」を返している
	chanOwner := func() <-chan int {
		// チャネルの所有者なので、チャネルを初期化している(定義1)
		intCh := make(chan int, 5)

		go func() {
			// チャネルの所有者なので、チャネルを閉じている(定義3)
			defer close(intCh)

			for i := 0; i <= 5; i++ {
				// チャネルの所有者なので、チャネルに書き込んでいる(定義2)
				intCh <- i
			}
		}()

		return intCh
	}

	// チャネルの所有者が、消費者にむけて「読み込んでよいぞチャネル」を作ってあげた感じ
	// このへん擬人法の怪しさがあふれる
	ch := chanOwner()

	// Mainゴルーチンは、チャネルの非所有者(チャネルを読み込むだけの人)
	for result := range ch {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Printf("Done receiving!")
}
