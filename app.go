package main

import (
	loop "github.com/shd00700/ModbusMQTT/Loop"
	"sync"
)
func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go loop.Loop(&wg)

	wg.Wait()
}
