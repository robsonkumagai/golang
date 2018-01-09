package repo

import mgo "gopkg.in/mgo.v2"

//SessaoMongo armazena a sessao de conexão com o mongo
var SessaoMongo *mgo.Session

//AbreSessaoComMongo abre a sessão do DB
func AbreSessaoComMongo() (err error) {
	err = nil
	SessaoMongo, err = mgo.Dial("mongodb://localhost:27017/cursodego")
	return
}
