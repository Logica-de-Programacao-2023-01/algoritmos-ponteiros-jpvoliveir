package main

import (
	"fmt"
	"math"
)

func updateWithPiAverage(num *float64) {
	pi := math.Pi                // Obtém o valor da constante Pi
	average := (*num + pi) / 2.0 // Calcula a média aritmética
	*num = average               // Atualiza o valor da variável para a média
}

func main() {
	number := 3.14
	fmt.Println("Antes:", number)
	updateWithPiAverage(&number)
	fmt.Println("Depois:", number)
}
