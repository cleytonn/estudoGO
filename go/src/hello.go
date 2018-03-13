package main //Pacote Principal

import "fmt" //Pacote de formatação, como: print e scanf

import "os"

func main() {

	exibeIntrodução()

	exibeMenu()

	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramento...")

	case 2:
		fmt.Println("Exibindo logs...")

	case 0:
		fmt.Println("Saindo do programa!")
		os.Exit(0)
	default:
		fmt.Println("Não conheco este comando")
		os.Exit(-1)
	}
}

func exibeIntrodução() {

	fmt.Println()

	nome := "Cleyton"
	versao := 1.1

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {

	fmt.Println("\n1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa\n")
}

func lerComando() int {

	var comando int

	fmt.Print("O que deseja fazer: ")
	fmt.Scan(&comando)

	return comando
}
