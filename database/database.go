package database

import (
	"database/sql"
	"os"

	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// NewDB function
func NewDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Erro ao carregar o arquivo .env: %v\n", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")


	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)

	if err != nil {

		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)

	}

	log.Println("Conectado ao banco de dados com sucesso!")

	return db, nil
}