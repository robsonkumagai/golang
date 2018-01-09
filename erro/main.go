package main

import (
	"encoding/json"
	"fmt"

	"github.com/robsonkumagai/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa do Robson"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(10056545646); err != nil {
		fmt.Println("[main] Houve um erro na atribuição	de valor a casa: ", err.Error())
		if err == model.ErrValorDeImovelAlto {
			fmt.Println("Corretor por favor verifique a sua avaliação")
		}
		return
	}

	fmt.Printf("A casa do Robson é: %+v\r\n", casa)
	objJSON, err := json.Marshal(casa)

	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto json: ", err.Error())
		return
	}

	fmt.Println("A casa em JSON: ", string(objJSON))

}
