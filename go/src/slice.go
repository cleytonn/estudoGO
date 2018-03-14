//Criando um slice através de uma função exibeNomes

package main

import (
	"fmt"
)

func main() {
	exibeNomes()
}

func exibeNomes() {
	nomes := []string{"João", "CLecio", "José"}

	//Imprimindo a quantidade de itens dentro do slice através da função len
	fmt.Println("O meu slice tem:", len(nomes), "itens")

	//Imprimindo a capacidade suportada suportada pelo slice através da função cap
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens\n")

	//Adiciona itens no slice, através da função append, que recebe o slice e o item a ser adicionado
	nomes = append(nomes, "Carlos")

	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
