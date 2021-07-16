package main

import "fmt"

// 疑問1: 複数のチャネルが準備OKな場合、selectはどうなるのか？
func main() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	// 速攻で、closeしている！ == チャネルが読み込める状態 == 準備OK状態
	close(c1)
	close(c2)

	var count1, count2 int

	for i := 1000; i > 0; i-- {
		select {
		case <-c1:
			count1++
		case <-c2:
			count2++
		}
	}

	// Q. 何が出力されるでしょう？
	fmt.Printf("count1: %d\ncount2: %d\n", count1, count2)
}
