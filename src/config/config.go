package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
)

// Carregar vai inicializar as variaveis de ambientes
func Carregar() {
	var erro error

	// Carrega variáveis do .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar .env: ", erro)
	}

	// Porta da API
	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		Porta = 9000 // valor padrão
	}

	// Monta string de conexão com todas variáveis
	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("MYSQLUSER"),
		os.Getenv("MYSQLPASSWORD"),
		os.Getenv("MYSQLHOST"),
		os.Getenv("MYSQLPORT"),
		os.Getenv("MYSQLDATABASE"),
	)
}
