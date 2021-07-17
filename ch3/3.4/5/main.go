package main

import (
	"fmt"
	"time"
)

// 疑問: チャネルが1つも読み込めず、その間に何かする必要がある場合にはどうしたらいいの？
func main() {
	start := time.Now()
	var c1, c2 <-chan int // どちらもnilチャネルなので永遠ブロック

	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after%v\n\n", time.Since(start))
	}
}
