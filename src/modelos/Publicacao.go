package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicacao feita por um usuario
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorID,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  string    `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

// Vai validar e formatar os campos
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O titulo não pode esta em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O conteudo é obrigatorio e não pode esta em branco")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
