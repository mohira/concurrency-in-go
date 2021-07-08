package main

import "fmt"

func main() {
	fmt.Println("mainゴルーチンはじまるよ〜")

	go sayHello()

	fmt.Println("mainゴルーチンおわるよ〜")
}

func sayHello() {
	fmt.Println("ゴルーチン2")
}
