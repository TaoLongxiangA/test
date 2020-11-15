package testInterface2

import "fmt"

type Cat struct {
}

func NewCat() *Cat {
	return &Cat{}
}

func (c *Cat) Picture() {
	fmt.Println("22222222")
}

func (c *Cat) Voice() {
	fmt.Println("2222")
}
