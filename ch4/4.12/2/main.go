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
		// cancelParent() を timeLimit より後に実行されるようにする
		time.Sleep(3 * time.Second)
		cancelParent()
		fmt.Println("親コンテキストでキャンセル！")
	}()

	// ctxChild.Done() がブロック解除されるパターンは2つ
	// 1. cancelParent() が Timeoutの timeLimit が「来る前」に実行された場合
	// 2. cancelParent() が Timeoutの timeLimit が「来た後」に実行された場合
	<-ctxChild.Done()

	fmt.Printf("ctxParent.Err() = %v\n", ctxParent.Err()) // <nil>
	fmt.Printf("ctxChild.Err() = %v\n", ctxChild.Err())   // context deadline exceeded

	fmt.Println("----- main() done ------")
}
