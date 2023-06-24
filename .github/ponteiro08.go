package main

import (
	"fmt"
)

func modifyVariable(ptr *int) {
	newValue := 42
	*ptr = newValue
}

func main() {
	value := 10
	fmt.Println("Antes:", value)

	modifyVariable(&value)

	fmt.Println("Depois:", value)
}
