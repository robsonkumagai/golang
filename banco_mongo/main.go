package main

import (
	"fmt"
	"net/http"

	"github.com/robsonkumagai/banco_sql/repo"
	"github.com/robsonkumagai/servidor_web/manipulador"
)

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Erro ao abrir banco de dados" err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local", manipulador.Local)
	

	fmt.Println("Servidor funcionando....")
	http.ListenAndServe(":8181", nil)
}
