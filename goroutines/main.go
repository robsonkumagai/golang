package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/robsonkumagai/goroutines/model"
)

func main() {

	traduzirParaJSON("saopaulo")
	traduzirParaJSON("riodejaneiro")
}

func traduzirParaJSON(nomeArquivo string) {
	fmt.Println(time.Now(), "- Começando a tradução do arquivo: ", nomeArquivo)
	arquivo, err := os.Open(nomeArquivo + ".csv")

	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}

	defer arquivo.Close()
	//Defer fecha o arquivo

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()

	if err != nil {
		fmt.Println("[main]	Houve um erro ao ler o arquivo com o leitor CSV. Erro: ", err.Error())
		return
	}

	arquivoJSON, err := os.Create("cidades.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
		return
	}

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Printf("Cidade: %+v\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)

			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o json do item ", item, "", err.Error())
			}

			escritor.WriteString(" " + string(cidadeJSON))
			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}

		}
	}

	escritor.WriteString("\r\n")
	escritor.Flush()
	fmt.Println(time.Now(), "- A tradução do arquivo: ", nomeArquivo, "foi finalizado", err.Error())
}
