package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := [2]string{
		"dog",
		"cat",
	}

	fmt.Println(animals)

	slice_animals := []string{
		"dog",
		"cat",
	}

	slice_animals = append(slice_animals, "bird")

	fmt.Println(slice_animals)

	slice_animals = slices.Delete(slice_animals, 0, 1) // delete in Go

	fmt.Println(slice_animals)

	// for loop
	for i := 0; i < len(slice_animals); i++ {
		fmt.Printf("The current animal is %s\n", slice_animals[i])
	}

	// for loop using
	for index, value := range slice_animals {
		fmt.Printf("current animal %s and current value %d", value, index)
	}

	// while loops
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

}
