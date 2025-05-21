package main

import "fmt"

/*
		Primitive data structures:
	    Strings
	    maps
	    arrays
	    slices
*/
func main21() {
	//x := 1
	//fmt.Println("Result: %i", x)

	var m = map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	//var j = make(map[string]int)
	for key, value := range m {
		fmt.Println("Result", key, value)
		//j[key] = value
	}

}
