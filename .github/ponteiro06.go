package main

import (
	"fmt"
)

type Livro struct {
	Titulo string
	Autor  string
}

func atualizarTitulo(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Aventura Misteriosa",
		Autor:  "Anônimo",
	}

	fmt.Println("Antes:", livro)

	atualizarTitulo(&livro)

	fmt.Println("Depois:", livro)
}
