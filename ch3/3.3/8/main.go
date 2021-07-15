package main

// nilチャネル
func main() {
	var ch chan int

	// <-ch // deadlock!
	// ch <- 1 // deadlock!
	close(ch) // panic: close of nil channel
}
