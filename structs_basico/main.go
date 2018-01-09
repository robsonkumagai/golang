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
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Ap Robson", 17514120}
	fmt.Printf("Meu apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Minha Chacara",
		valor: 1120,
		X:     22,
	}
	fmt.Printf("A chácara é: %+v\r\n", chacara)

	casa.Nome = "Minha casa"
	casa.valor = 680000
	casa.X = 18
	casa.Y = 31

	fmt.Printf("A nova casa é: %+v\r\n", casa)
}
