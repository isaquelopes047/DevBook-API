package router

import "github.com/gorilla/mux"

// Gerar o router e retornar as rotas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
