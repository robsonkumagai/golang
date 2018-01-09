package main

import (
	"fmt"

	"github.com/robsonkumagai/cursodego/variaveis/pacotes/operadora"
	"github.com/robsonkumagai/cursodego/variaveis/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Robson"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}
