package main

import "github.com/robsonkumagai/structs_avancado/model"
import "fmt"
import "encoding/json"

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa do Robson"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(75800)

	fmt.Printf("A casa do Robson é: %+v\r\n", casa)
	fmt.Printf("O valor da casa é: %d\r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)

	/* _ Ignora o retorno
	 * String é um array de bits
	 */

	fmt.Println("A casa em JSON: ", string(objJSON))
}
