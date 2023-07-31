package main

import "fmt"

func isPair(number int64) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	result := isPair(10)
	fmt.Println(result)
}

/* func main() {
	for a := 1; a <= 10; a++ {
		if a % 2 == 0 {
			fmt.Printf("O número %d é par\n", a)
		} else {
			fmt.Printf("O número %d é ímpar\n", a)
		}
	}
}
 */