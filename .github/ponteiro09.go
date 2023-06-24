package main

import (
	"fmt"
)

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livro) {
	desconto := livro.Preco * 0.1
	livro.Preco -= desconto
}

func main() {
	livro := Livro{
		Titulo: "Aventura Misteriosa",
		Autor:  "John Doe",
		Preco:  50.0,
	}

	fmt.Println("Preço antes do desconto:", livro.Preco)

	aplicarDesconto(&livro)

	fmt.Println("Preço após o desconto:", livro.Preco)
}
