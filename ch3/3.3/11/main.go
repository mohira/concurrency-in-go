package main

// 双方向 => 受信専用 は OK
func to1(c chan int) <-chan int {
	return c
}

// 双方向 => 送信専用 は OK
func to2(c chan int) chan<- int {
	return c
}

// 一方向 => 双方向 は無理
// cannot use c (type <-chan int) as type chan int in return argument
func to3(c <-chan int) chan int {
	return c
}

// 一方向 => 双方向 は無理
// cannot use c (type chan<- int) as type chan int in return argument
func to4(c chan<- int) chan int {
	return c
}
