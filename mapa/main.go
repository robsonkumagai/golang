package main

import "github.com/robsonkumagai/mapa/model"
import "fmt"

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa do Robson"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(75800)

	cache["Casa do Robson"] = casa

	fmt.Println("A casa do Robson está no cache?")
	imovel, achou := cache["Casa do Robson"]
	if achou {
		fmt.Printf("Sim, achei a casa: %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto da Ju"
	apto.X = 29
	apto.Y = 45
	apto.SetValor(95800)

	cache[apto.Nome] = apto

	//Quantidade de itens que está no cache
	fmt.Println("Quantos itens há no cache?", len(cache))

	//Mostrando todos os valores que estão no cache
	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa do Robson"]
	if achou {
		delete(cache, imovel.Nome)
	}
	//Deleta do cache
}
