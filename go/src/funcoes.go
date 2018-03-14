//Neste código acontecem alguns retornos dentro da mesma função, como string e int.

package main

import "fmt"

func main() {

	nome, idade := devolveNomeIdade()

	fmt.Println("Nome", nome, ", idade:", idade)
}

func devolveNomeIdade() (string, int) {

	nome := "Cleyton"

	idade := 18

	return nome, idade
}
