package main

func receiveOnlyChannel() <-chan int {
	c := make(chan int) // この時点では読み書きできるチャネル

	return c
}

func main() {
	// 返り値の型のおかげで受信専用のチャネルになっている
	// 呼び出し側が書き込むことを制限できている感じ
	ch := receiveOnlyChannel()

	// 許されていない行為(書き込み)をしようとしてもできなくなっている(嬉しい)
	ch <- 1 // invalid operation: c <- 1 (send to receive-only type <-chan int)
}
