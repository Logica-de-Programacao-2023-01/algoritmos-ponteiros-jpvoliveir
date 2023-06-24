package main

import (
	"fmt"
)

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValor(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	conta := Conta{
		Saldo:   1000.0,
		Titular: "João",
	}

	fmt.Println("Saldo antes:", conta.Saldo)

	adicionarValor(&conta, 500.0)

	fmt.Println("Saldo depois:", conta.Saldo)
}
