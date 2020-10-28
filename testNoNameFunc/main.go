package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(begin time.Time) {
		fmt.Println(time.Since(begin))
	}(time.Now())

	time.Sleep(20 * time.Second)
}
