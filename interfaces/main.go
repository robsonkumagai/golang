package main

import "github.com/robsonkumagai/interfaces/model"
import "fmt"

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"
}

func acordarComCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func ouvirPatoNoLago(p model.Pato) {
	fmt.Println(p.Quack())
}
