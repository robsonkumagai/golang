package model

//Cidade representa a cidade e o estado do Brasil
type Cidade struct {
	Nome   string `json:"nome"`
	Estado string `json:"estado"`
}
