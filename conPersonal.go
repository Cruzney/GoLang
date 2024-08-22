package main

import (
	"fmt"
	"time"
)

const alongamento = 1
const aquecimento = 5
const corridaLeve = 10
const corridaForte = 10

func main() {
	fmt.Println("Período de alongamento...")
	time.Sleep(alongamento * time.Second)

	fmt.Println("Período de aquecimento...")
	time.Sleep(aquecimento * time.Second)

	fmt.Println("Período de corrida leve...")
	time.Sleep(corridaLeve * time.Second)

	fmt.Println("Período de corrida forte...")
	time.Sleep(corridaForte * time.Second)

	fmt.Println("Período de alongamento...")
	time.Sleep(alongamento * time.Second)
}
