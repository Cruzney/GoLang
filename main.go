package main

import (
	"fmt"
	"reflect"
)

func main() {
	var greeting string = "Hello"
	name := "Edinei Cruz"
	versao := 1.1 // desta forma posso encurtar o código usando ":=" e já deixando o Go detectar o tipo
	var titulo string = "Sr./Sra."
	var msgPrograma string = "Este programa está na versão:"

	message := greeting + " " + titulo + " " + name

	fmt.Println(message)
	fmt.Println(msgPrograma, versao)
	fmt.Println("A variável versão e do tipo:", reflect.TypeOf(versao))

	//var comando int
	var comandoo int
	fmt.Scan(&comandoo) // função que não precisa de endereço ou modificador
	//fmt.Scanf("%d", &comando) // função que precisa do endereço fisico
	fmt.Println("O Enderço da minha variavél comando é,", &comandoo)
	fmt.Println("O Comando Escolhido foi:", comandoo)
}
