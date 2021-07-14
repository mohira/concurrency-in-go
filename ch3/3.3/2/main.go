package main

func main() {
	sendOnlyChannel := make(chan<- int)
	receiveOnlyChanel := make(<-chan int)

	// invalid operation: <-sendOnlyChannel (receive from send-only type chan<- int)
	<-sendOnlyChannel // 受信しようとしてみる

	// invalid operation: receiveOnlyChanel <- 1 (send to receive-only type <-chan int)
	receiveOnlyChanel <- 1 // 送信しようとしてみる

}
