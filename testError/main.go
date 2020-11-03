package main

import (
	"fmt"
	"log"
)

func main() {
	err := testError()
	if err != nil {
		fmt.Println(err)
	}
}

func testError() (err error) {
	log.Println(err)
	return err
}
