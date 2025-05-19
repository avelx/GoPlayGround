package main

import "fmt"

func main() {
	// Idiomatic:
	v := make([]int, 2, 5)

	fmt.Println("This is a result: %i", v)
}
