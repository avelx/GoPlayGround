package main

import (
	"fmt"
	"log"
	"os"
)

// Sum Pass a pointer of an array
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

//array := [...]float64{7.0, 8.5, 9.1}
//x := Sum(&array)  // Note the explicit address-of operator

func main() {
	/*
		Arrays are values. Assigning one array to another copies all the elements.
		In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
		The size of an array is part of its type. The types [10]int and [20]int are distinct.
	*/

	array := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	fmt.Println("Result: %i", x)

	// Slice via a => read file

	file, _ := os.Open("data/test.dat") // For read access.

	data := make([]byte, 100)
	var count, err = file.Read(data)
	for count > 0 {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
		count, err = file.Read(data)
	}

}
