package repo

import (
	"github.com/jmoiron/sqlx"
	/*__ "github.com/go-sql-driver/mysql" não é usado diretamente pela aplicação*/)

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL é função que abre a conexão com o BD mysql
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/cursodego")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return

}
