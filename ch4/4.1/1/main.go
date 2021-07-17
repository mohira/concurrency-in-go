package main

import "fmt"

func main() {
	data := make([]int, 4) // こいつが拘束対象の囚人か？

	// やろうと思えば、独房の外で囚人を操作できてしまう！
	// 規約違反！
	// data[3] = 333
	// data[4] = 444

	loopData := func(handleData chan<- int) {
		// loopData という独房
		// ここにdataという囚人を拘束する規約になっている ← アドホック！
		defer close(handleData)

		// 独房内で囚人を操作できる(これは独房内だから規約に準じている)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
