package main

import (
	"fmt"
	"math/rand"
)

// p.94 ゴルーチンリーク
// ゴルーチンがチャネルに対して書き込みを行おうとしてブロックしている状況
func main() {
	newRanStream := func() <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandomStream closure exited.")
			defer close(randStream)

			for {
				// 終了条件がないので、ずーっと書き込みを頑張り続ける
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRanStream()
	fmt.Println("3 random ints:")

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
