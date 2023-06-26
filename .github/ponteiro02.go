package main

import "fmt"

func checkParImpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0 // Par
	} else {
		*ptr = 1 // Ímpar
	}
}

func main() {
	num := 7

	fmt.Printf("Antes: %d\n", num)

	checkParImpar(&num)

	fmt.Printf("Depois: %d\n", num)
}
