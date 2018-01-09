package main

import "fmt"

/*
* Quando a variável iniciar com letra maiúscula ela é pública
* Quando iniciar com minúscula é privada
* Quando a variável for pública, fazer o comentário para definir o valor dela
 */

var (
	endereco   string //endereco = ""
	telefone   = "99700-2802"
	quantidade int     //quantidade = 0
	comprou    bool    //comprou = false
	valor      float64 //valor = 0.00
	palavras   rune
)

func main() {
	teste := "Valor do teste" //Define a variável e já coloca o valor

	fmt.Printf("Endereco: %s\r\n", endereco)
	fmt.Printf("Quantidade: %d\r\n", quantidade)
	fmt.Printf("Comprou: %v\r\n", comprou)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
