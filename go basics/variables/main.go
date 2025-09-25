package main

import (
	"fmt"
)

func main() {
	// var myName string = "Aditya"
	myName := "Aditya"
	myInt := 10
	myFloat := 10.0

	fmt.Printf("Hello my name is %s my int is %d and my float is %f\n", myName, myInt, myFloat)

	var myFriendsName string // Go has a concept of "zero-values" for each type
	var myBool bool
	var myOtherInt int

	fmt.Printf("my other friend %s my bool %t and my int is %d\n", myFriendsName, myBool, myOtherInt)
}
