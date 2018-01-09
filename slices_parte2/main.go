package main

import (
	"fmt"
)

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Brasilia")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 5)
	cidades[0] = "New York"
	cidades[1] = "Londres"
	cidades[2] = "Tokyo"
	cidades[3] = "Singapura"
	cidades[4] = "Seol"

	fmt.Println(cidades, len(cidades), cap(cidades))

	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}

	//Primeiro item começa com indice zero
	//Segundo item começa com indice um
	capitaisAsia := cidades[3:5]
	fmt.Println(capitaisAsia)

	temp1 := cidades[:2]
	fmt.Println(temp1)

	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)

	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	copy(cidades, temp)
	fmt.Println(cidades)

}
