package main

import "fmt"

const ebulitionF = 212.0

func main() {
	F := ebulitionF
	C := (F - 32) * 5 / 9
	fmt.Printf("O ponto de ebulição da água em F é: %g, e em C é: %g ", F, C)
}