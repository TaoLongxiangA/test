package main

import (
	"fmt"
	"log"
)

func main() {
	var a interface{}
	test0(a)

	//fmt.Println(f)
}

func test0(a interface{}) {
	//return func(a interface{}) bool {
	i := a.(int)
	log.Println("convert success", i)
	//}
}

func test1() {
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
