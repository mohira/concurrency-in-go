package main

func main() {
	ch := make(chan int)

	// ずーっと受信を待つゴルーチン
	go func() {
		<-ch
	}()
}
