package main

import "fmt"

func main9() {
	// Idiomatic:
	// applicable only to => maps, slices and channels
	v := make([]int, 2, 5)

	fmt.Println("This is a result: %i", v)
}
