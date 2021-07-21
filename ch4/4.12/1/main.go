package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxParent, cancelParent := context.WithCancel(context.Background())

	timeLimit := 1 * time.Second
	ctxChild, cancelChild := context.WithTimeout(ctxParent, timeLimit)
	defer cancelChild()

	go func() {
		// cancelParent() を timeLimit より先に実行されるようにする
		time.Sleep(500 * time.Millisecond)
		cancelParent()
		fmt.Println("親コンテキストでキャンセル！")
	}()

	// ctxChild.Done() がブロック解除されるパターンは2つ
	// 1. cancelParent() が Timeoutの timeLimit が「来る前」に実行された場合
	// 2. cancelParent() が Timeoutの timeLimit が「来た後」に実行された場合
	<-ctxChild.Done()

	// どちらも同じ値になる!
	// 親コンテキスト(WithCancel)がキャンセルされたので、
	// 子コンテキスト(WithTimeout)もキャンセルされる
	fmt.Printf("ctxParent.Err() = %v\n", ctxParent.Err()) //  context canceled
	fmt.Printf("ctxChild.Err() = %v\n", ctxChild.Err())   //  context canceled

	fmt.Println("----- main() done ------")
}
