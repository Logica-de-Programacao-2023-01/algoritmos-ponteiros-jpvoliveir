package main

import (
	"fmt"
)

func sumLastDigits(num *int) {
	digit1 := *num % 10        // Obtém o último dígito
	digit2 := (*num / 10) % 10 // Obtém o penúltimo dígito
	sum := digit1 + digit2     // Calcula a soma dos dois últimos dígitos
	*num = sum                 // Atualiza o valor da variável para a soma
}

func main() {
	number := 1234
	fmt.Println("Antes:", number)
	sumLastDigits(&number)
	fmt.Println("Depois:", number)
}
