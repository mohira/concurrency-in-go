package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	// printData以前に宣言するとやりたい放題で大変！
	data := []byte("golang")
	var buf bytes.Buffer

	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		for _, b := range data {
			fmt.Fprintf(&buf, "%c", b)
		}

		fmt.Println(buf.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
