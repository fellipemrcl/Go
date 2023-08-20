package main

import "fmt"

func main () {
	var x [4] float64 
	x [0] = 5.5
	x [1] = 8
	x [2] = 5
	x [3] = 7
	var total float64 = 0
	for i := 0; i < len(x); i ++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}