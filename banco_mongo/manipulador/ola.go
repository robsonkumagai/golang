package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/robsonkumagai/servidor_web/model"
)

//Ola é um manipulador da requisição a rota /ola
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := Modelos.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucão do modelo: ", err.Error())
	}
}
