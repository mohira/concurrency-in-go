package main

func main() {
	c1 := make(<-chan interface{})
	c2 := make(<-chan interface{})
	c3 := make(chan<- interface{})

	select {
	case <-c1:
	case <-c2:
	case c3 <- struct{}{}:
	}
}
