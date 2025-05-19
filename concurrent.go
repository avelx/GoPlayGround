package main

import (
	"fmt"
	"runtime"
)

type Vector []float64

//const numCPU = 4 // number of CPU cores

func (v Vector) DoAll(u Vector) {
	var numCPU = runtime.NumCPU()
	fmt.Println("Number of CPUs:", numCPU)
	c := make(chan int, numCPU) // Buffering optional but sensible.
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}
	// Drain the channel.
	for i := 0; i < numCPU; i++ {
		<-c // wait for one task to complete
	}
	// All done.
}

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	v[i] += u[i] + 1
	c <- 1 // signal that this piece is done
}

func main4() {
	var v = Vector{1.45, 5.45, 8.19, 2.91}
	var j = Vector{8.45, 3.45, 16.19, 78.91}
	v.DoAll(j)

	fmt.Println("This is a result: %i", v)
}
