package main

import (
	"fmt"
)

//type Post struct {
//	ID      int
//	Content string
//	Author  string
//}

func main() {
	//post1 := Post{
	//	ID:      1,
	//	Content: "Hello World!",
	//	Author:  "Sau Sheong",
	//}

	intArray := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(intArray), cap(intArray))

	t := make([]int, len(intArray), (cap(intArray)+1)*2)
	copy(t, intArray)

	intArray = t

	fmt.Println("len and cap currency")
	fmt.Println(len(intArray), cap(intArray))

	a := make([]int, 1)
	fmt.Println(len(a), cap(a))
	a = append(a, 2, 3, 4, 5)
	fmt.Println(len(a), cap(a))

}
