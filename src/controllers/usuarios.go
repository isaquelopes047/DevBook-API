package controllers

import (
	"net/http"
)

// CriarUsuario registra um usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario"))
}

// BuscarUsuarios busca usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios!"))
}

// BuscarUsuario busca um usuario
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuarios!"))
}

// AtualizarUsuario atualiza o registro
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuarios!"))
}

// DeletarUsuario deelta o registro
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuarios!"))
}
