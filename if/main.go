package main

import (
	"fmt"
)

func main() {
	meses := 0
	situacao := true
	cidade := "Marília"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Esse credor é inadimplente")
	}

	if !situacao {
		fmt.Println("Esse credor é adimplente")
	}

	if cidade == "Marília" {
		fmt.Printf("O cliente vive na cidade %s\r\n", cidade)
	}

	if descricao, status := tempoDevendo(meses); status {
		fmt.Println("Qual a situação do cliente: ", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")
	//Caso entre no if acima não será mostrado está mensagem
}

func tempoDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente é inadimplente"
		return
	}

	descricao = "O cliente está em dia"
	return
}
