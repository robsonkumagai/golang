package model

//Imovel representa as informações de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor seta o valor do imóvel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor
func (i *Imovel) GetValor() int {
	return i.valor
}
