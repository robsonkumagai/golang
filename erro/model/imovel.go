package model

import "errors"

//Imovel representa as informações de um imóvel
type Imovel struct {
	Nome  string `json:"nome"`
	valor int
	X     int `json:"coordenada_x"`
	Y     int `json:"coordenada_y"`
}

//ErrValorDeImovelInvalido é um erro que é mostrado quando o imóvel está com o valor abaixo do mercado
var ErrValorDeImovelInvalido = errors.New("O valor informado do imóvel é inválido")

//ErrValorDeImovelAlto é um erro que é mostrado quando o imóvel está com o valor acima do mercado
var ErrValorDeImovelAlto = errors.New("O valor informado do imóvel é inválido")

//GetValor retorna o valor
func (i *Imovel) GetValor() int {
	return i.valor
}

//SetValor seta o valor do imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 1000000 {
		err = ErrValorDeImovelAlto
		return
	}
	i.valor = valor
	return
}
