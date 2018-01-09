package model

import (
	"database/sql"
)

//Local armazena os dados da localidade pelo seu código telefonico
type Local struct {
	Pais             string         `json:"pais" db:"country" bson:"country"`
	Cidade           sql.NullString `json:"cidade" db:"city" bson:"city"`
	CodigoTelefonico int            `json:"cod_telefone" db:"telcode" bson:"telcode"`
}
