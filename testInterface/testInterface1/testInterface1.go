package testInterface1

import (
	"fmt"
	"log"
	"test/testInterface/testInterface2"
)

type Animal interface {
	Picture()
	Voice()
}

type Dog struct {
}

func (d *Dog) Picture() {
	log.Println("11111111")
	cat := testInterface2.NewCat()
	cat.Picture()
	fmt.Println("yes")
}

func (d *Dog) Voice() {
	log.Println("1111")
}
