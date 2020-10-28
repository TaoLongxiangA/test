package main

import "fmt"

func main() {
	//go func() {
	a := func(i interface{}) string {
		if i != nil {
			return "111"
		}
		return "222"
	}
	//}()
	fmt.Printf("%d ", a)
}
