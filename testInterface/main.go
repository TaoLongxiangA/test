package main

import "test/testInterface/testInterface1"

func main() {
	var dog *testInterface1.Dog
	test(dog)
}

func test(s testInterface1.Animal) {
	s.Picture()
}
