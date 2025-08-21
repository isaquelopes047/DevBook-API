package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	email := strings.TrimSpace(usuario.Email)
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioSalvoNoBanco.ID == 0 || usuarioSalvoNoBanco.Senha == "" {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("e-mail e/ou senha inválidos"))
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("e-mail e/ou senha inválidos"))
		return
	}

	token, erro := autenticacao.CriarToken(uint64(usuarioSalvoNoBanco.ID))
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	fmt.Println((token))
	w.Write([]byte(token))
}
