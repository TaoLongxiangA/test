package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	log.Println(1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//errChan := make(chan error)
	log.Println(2)
	errChan := <-c //block
	log.Println(3)
	fmt.Println(errChan)
}
