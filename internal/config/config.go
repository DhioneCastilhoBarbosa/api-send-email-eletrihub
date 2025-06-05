package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado. Usando variáveis de ambiente.")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
