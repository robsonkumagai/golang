package main

import (
	"fmt"
	"net/http"

	"github.com/robsonkumagai/servidor_web/manipulador"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)

	fmt.Println("Servidor funcionando....")
	http.ListenAndServe(":8181", nil)
}
