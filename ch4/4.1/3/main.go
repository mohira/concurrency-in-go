package main

import (
	"bytes"
	"fmt"
	"sync"
)

// buffの定義を広くすると壊れる例もあげておこう

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		// printData宣言以前にdataが宣言されてないので
		// コピーを受け取るしかない
		// printDataを起動するゴルーチンがどう渡すかに制限されるわけ
		defer wg.Done()

		var buf bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buf, "%c", b)
		}

		fmt.Println(buf.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	data := []byte("golang")

	// printDataが使えるdataを一部の範囲に絞っている == 拘束
	go printData(&wg, data[:3]) // 先頭3バイトだけ操作できる拘束
	go printData(&wg, data[3:]) // 後半3バイトだけ操作できる拘束

	wg.Wait()
}
