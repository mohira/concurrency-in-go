package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)

		go func(salutation string) {
			fmt.Println(salutation)
			wg.Done()
		}(salutation)
	}

	wg.Wait()
}
