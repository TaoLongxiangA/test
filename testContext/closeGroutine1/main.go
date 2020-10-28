package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for true {
			select {
			case <-stop:
				fmt.Println("goroutine stop...")
				return
			default:
				fmt.Println("watching...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("stop watching")
	stop <- true
	time.Sleep(5 * time.Second)

}
