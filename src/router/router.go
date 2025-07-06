package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar o router e retornar as rotas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
