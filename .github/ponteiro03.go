package main

import (
	"fmt"
)

func reverseString(str *string) {
	runes := []rune(*str) // Converte a string em um slice de runes
	length := len(runes)  // Obtém o comprimento da string
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Inverte os caracteres
	}
	*str = string(runes) // Atualiza o valor da string reversa
}

func main() {
	text := "Olá, mundo!"
	fmt.Println("Antes:", text)
	reverseString(&text)
	fmt.Println("Depois:", text)
}
