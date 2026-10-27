package database

import (
	"log"
	"net"
	"net/url"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectaComBancoDeDados() {
	stringDeConexao := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(getenv("DB_USER", "postgres"), getenv("DB_PASSWORD", "postgres")),
		Host:   net.JoinHostPort(getenv("DB_HOST", "localhost"), getenv("DB_PORT", "5432")),
		Path:   "/" + getenv("DB_NAME", "root"),
	}
	query := stringDeConexao.Query()
	query.Set("sslmode", "disable")
	stringDeConexao.RawQuery = query.Encode()

	var err error
	DB, err = gorm.Open(postgres.Open(stringDeConexao.String()))
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados: %v", err)
	}

	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Panicf("Erro ao migrar banco de dados: %v", err)
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
