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
	casa.SetValor(100)

	fmt.Printf("A casa do Robson é: %+v\r\n", casa)
	objJSON, _ := json.Marshal(casa)

	fmt.Println("A casa em JSON: ", string(objJSON))
}
