package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Print("O número ", numero, "se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	fmt.Print("Você utiliza Unix? ")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim")
	default:
		fmt.Println("Não")
	}

	//É possível utiizar duas condições para o switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ok, você pode dormir até mais tarde!")
	default:
		fmt.Println("Bora trabalhar!")
	}

	numero = 10
	fmt.Println("Este número cabe num dígito?")
	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Dois digitos")
	case numero >= 100:
		fmt.Println("Número muito grande")
	}

}
