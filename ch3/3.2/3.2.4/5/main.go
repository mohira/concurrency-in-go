package main

import (
	"sync"
)

// sync.Mutex をコピーする
// コンパイルは通るが、警告がでる
func main() {
	var m1 sync.Mutex
	var m2 sync.Mutex
	m2 = m1 // assignment copies lock value to m2: sync.Mutex
	_ = m2  // assignment copies lock value to _: sync.Mutex
}
