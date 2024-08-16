package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			startMonitoring()
			// iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço este comando...")
			os.Exit(-1)
		}

	}

}

func exibeIntroducao() {
	var greeting string = "Hello"
	name := "Edinei Cruz"
	versao := 1.1                  // desta forma posso encurtar o código usando ":=" e já deixando o Go detectar o tipo
	var titulo string = "Sr./Sra." // Neste caso não preciso de ":=" para fazer o encurtamento
	var msgPrograma string = "Este programa está na versão:"

	message := greeting + " " + titulo + " " + name

	fmt.Println(message)
	fmt.Println(msgPrograma, versao)
	fmt.Println("A variável versão e do tipo:", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println(" 1 - Iniciar Monitormanento")
	fmt.Println(" 2 - Exibir Logs")
	fmt.Println(" O Sair do Programa")
}

func leComando() int {
	//var comando int
	var comandoLido int
	fmt.Scan(&comandoLido) // função que não precisa de endereço ou modificador
	//fmt.Scanf("%d", &comando) // função que precisa do endereço fisico
	fmt.Println("O Enderço da minha variavél comando é,", &comandoLido)
	fmt.Println("O Comando Escolhido foi:", comandoLido)

	return comandoLido
}

func startMonitoring() {
	fmt.Println("Iniciando Monitoramento...")

	site := "https://www.totvs.com.br"
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("O site ## ", site, " ## foi carregado com sucesso!")
	} else {
		fmt.Println("Ocorreu um erro: Status Code:", resp.StatusCode)
	}
}

// func iniciarMonitoramento() {
// 	fmt.Println("Monitorando...")
// 	// site com URL inexistente
// 	site := "https://www.alura.com.br" // ou 200
// 	resp, _ := http.Get(site)

// 	if resp.StatusCode == 200 {
// 		fmt.Println("Site:", site, "foi carregado com sucesso!")
// 	} else {
// 		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
// 	}
// }
