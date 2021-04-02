package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World")

	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint)

	var men uint8
	men = 5
	var women uint16
	women = 6

	var people int

//	people = men + women

	people = int(men) + int(women) //casting data type

	fmt.Println(people)
}
