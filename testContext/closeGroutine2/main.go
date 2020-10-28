package main

import (
	"fmt"
	"time"
)

func main() {
	signal := make(chan bool)

	for i := 0; i < 5; i++ {
		go testCloseGoroutine(signal, i)
	}

	time.Sleep(5 * time.Second)
	close(signal)
	time.Sleep(5 * time.Second)

}

func testCloseGoroutine(ch chan bool, i int) {
	for true {
		select {
		case v := <-ch:
			fmt.Printf("stop watching : %v , number is %v\n", i, v)
			return
		default:
			fmt.Printf("watching %d...\n", i)
			time.Sleep(2 * time.Second)
		}
	}
}
