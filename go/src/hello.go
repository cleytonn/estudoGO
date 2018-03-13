package main //Pacote Principal

import "fmt" //Pacote de formatação, como: print e scanf

func main() {
	nome := "Cleyton"
	versao := 1.1

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("\n1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa\n")

	var comando int

	fmt.Scan(&comando)

	fmt.Println("\nO valor da variável comando é:", comando)

}
