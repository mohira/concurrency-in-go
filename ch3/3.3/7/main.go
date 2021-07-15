package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	ch := make(chan int, 4) // バッファサイズ: 4

	go func() {
		defer close(ch)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

		// バッファサイズ分だけ整数を送信する(事前にわかっている)
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "っ'-')╮ =͞  Sending: %d\n", i)
			ch <- i
		}
	}()

	for integer := range ch {
		fmt.Fprintf(&stdoutBuff, "∠( ﾟдﾟ)／ Received %v.\n", integer)
	}
}
