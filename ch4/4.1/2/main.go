package main

import "fmt"

func main() {
	// 読み込み専用チャネルを返す時点で、所有者は「他の連中(ゴルーチン)には読み込みしか許可しないぞ！」って意思表示してる感じ
	chanOwner := func() <-chan int {
		// resultsをこの関数のレキシカルスコープ内で初期化している
		// → 操作できるのはこの関数(独房)の中だけ
		// → 書き込み権限の「拘束」ってわけ
		results := make(chan int, 5) // 今回の囚人

		go func() {
			// 所定の書き込みが終わったらclose(他の人は書き込めない)
			// まさに、チャネルの所有者の行いですな
			defer close(results)

			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()

		return results
	}

	// 拘束を突破しようとしても....
	// 独房の外からは、囚人に書き込みができない！(嬉しい)
	// results <- 999 // コンパイルエラー！

	consumer := func(results <-chan int) {
		// resultsの読み込み権限"だけ"のコピーを受け取っている

		// 拘束を突破しようとしても....
		// results <- 999 // コンパイルエラー！(嬉しい)

		// 約束通りの読み込み作業
		for result := range results {
			fmt.Printf("Received %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}
