package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int

	select {
	case <-c: // nilチャネルなので永遠にブロック
	case <-time.After(5 * time.Second): // time.Afterは読み込み専用チャネルを返している
		fmt.Println("Timed out")
	}
}
