package main

import "fmt"

var A int = 20
var B string = "Fellipe"

func main() {
	var C float64 = float64(A)
	fmt.Printf("O valor de C é: %g e o valor de B é: %s", C, B)
}