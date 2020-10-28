package main

import (
	"fmt"
	"reflect"
)

func main() {
	of := reflect.TypeOf(cat{})
	valueOf := reflect.ValueOf(cat{})
	name := reflect.Indirect(valueOf).Type().Name()
	fmt.Println(of, "\n", valueOf, "\n", name)
}

type speak interface {
	talk() string
}

type cat struct {
	fat string
}

func (c *cat) talk() string {
	return "2333"
}
