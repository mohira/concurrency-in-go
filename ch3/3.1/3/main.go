package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello" // salutationは「あいさつ」とか「敬礼」とかって意味らしい

	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
		fmt.Println(salutation, &salutation)
	}()

	wg.Wait()
	fmt.Println(salutation, &salutation)
}
