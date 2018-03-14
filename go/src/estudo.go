package main //Pacote Principal

import (
	"fmt" //Pacote de formatação, como: print e scanf
	"net/http"

	//Responsável pela requisição do sistema (get e post), fazendo a
	//comunicação com a web

	"os" //Trabalha com saídas do sistema
)

func main() {

	exibeIntrodução()

	for {

		exibeMenu() //Menu de opções

		comando := lerComando()

		switch comando {
		case 1:
			iniciaMonitoramento()

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

//No código abaixo o Operador diz para o GO ignorar o segundo retorno através do undescore
//Isto acontece porque o http.get(site) esta retornando mais de uma resposta

func iniciaMonitoramento() {
	fmt.Println("Iniciando monitoramento...")

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
	}
}
