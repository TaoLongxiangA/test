package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	str := "0x777"
	a, _ := strconv.Atoi(str) //son/a part of ParseInt
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)

	fmt.Println("------------------")

	b, err := strconv.ParseInt(str, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(int(b))
}
