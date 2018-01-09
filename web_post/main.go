package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/robsonkumagai/web_post/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 4
	usuario.Nome = "Robson Kumagai"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o JSON do objeto usuario", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://requestb.in/1284vor1", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request bin. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)

	if err != nil {
		fmt.Println("[main] Erro ao chamar o serviço do request bin. Erro: ",
			err.Error())
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo retornado pelo request bin. Erro: ",
				err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
