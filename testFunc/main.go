package main

import (
	"fmt"
	"time"
)

func main() {
	func(begin time.Time) {
		time.Sleep(5 * time.Second)
		fmt.Println(time.Since(begin))
	}(time.Now())
}
