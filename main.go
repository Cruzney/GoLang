package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const intervalo = 5
const logFilePath = "log.txt"
const sitesFilePath = "sites.txt"

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
			exibeLogs()
			// Exibindo logs
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
	//fmt.Println("A variável versão e do tipo:", reflect.TypeOf(versao))
}

func exibeMenu() {
	pulaLina()
	fmt.Println(" 1 - Iniciar Monitormanento")
	fmt.Println(" 2 - Exibir Logs")
	fmt.Println(" 0 - Sair do Programa")
	pulaLina()
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

	sites := readSitesInFile()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			TestSites(site)
		}
		time.Sleep(time.Second * intervalo)
		pulaLina()
	}

	pulaLina()

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

//	func treinaFor() {
//		pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21, 40}
//		for i, ponto := range pontosPlanningPoker {
//			fmt.Println(" o ponto na posição", i, " tem o valor:", ponto)
//		}
//	}
func TestSites(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}

	//defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("O site ## ", site, " ## foi carregado com sucesso!")
		registaLog(site, true)
	} else {
		fmt.Println("Ocorreu um erro: Status Code:", resp.StatusCode, " do site:", site)
		registaLog(site, false)
	}
}

func readSitesInFile() []string {

	var sites []string

	arquivo, err := os.Open(sitesFilePath)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registaLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //0666 = permissoes de escrita e leitura

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n") // 02/01/2006 15:04:05 = formato da data conforme documentação do Go

	arquivo.Close()
}

func exibeLogs() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := os.ReadFile(logFilePath)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

func pulaLina() {
	fmt.Println("")
}
