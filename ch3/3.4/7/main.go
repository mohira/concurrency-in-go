package main

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	//
	// goroutine 1 [select (no cases)]:
	select {}
}
