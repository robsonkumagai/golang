package prefixo

import "strconv"

//Capital representa o número do prefixo de telefone da capital de um estado
var Capital = 11

var teste = "Teste"

//TesteComPrefixo é uma variável pública de teste
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
