package main

import "fmt"

func main() {
	end := testReturn(100)(100)
	fmt.Println(end)
}

func testReturn(i int) End {
	return func(j int) int {
		return i + j
	}
}

type End func(int) int
