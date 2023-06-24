package main

import "fmt"

func sumOfFirstN(ptr *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*ptr = sum
}

func main() {
	var result int
	n := 5

	sumOfFirstN(&result, n)

	fmt.Printf("A soma dos primeiros %d números naturais é: %d\n", n, result)
}
