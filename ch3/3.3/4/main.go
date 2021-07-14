package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello channels!"
	}()

	salutation, ok := <-ch
	fmt.Printf("(%v): %v", ok, salutation)
}
