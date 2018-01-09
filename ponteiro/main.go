package main

import (
	"fmt"
)

//Imovel é uma struct que armazena os dados de um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel)
	fmt.Printf("Casa é: %v\r\n", casa)
	fmt.Printf("Endereço de memória da Casa é: %p + %v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara do Jeff", 28000}
	fmt.Printf("A chacara é: %p - %v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("A chacara é: %p - %v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
