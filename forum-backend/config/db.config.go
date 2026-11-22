package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	connectionString := os.Getenv("DATABASE_URL")

	if connectionString == "" {
		user := GetRequiredEnv("DB_USER")
		psswd := GetEnvWithDefault("DB_PWD", "")
		host := GetRequiredEnv("DB_HOST")
		port := GetRequiredEnv("DB_PORT")
		name := GetRequiredEnv("DB_NAME")

		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, psswd, host, port, name)
	}

	dbContext, dbContextErr := sql.Open("mysql", connectionString)
	if dbContextErr != nil {
		log.Fatalf("Erreur connection base de données - Erreur : \n\t%s", dbContextErr.Error())
	}

	pingErr := dbContext.Ping()
	if pingErr != nil {
		dbContext.Close()
		log.Fatalf("Erreur ping base de données - Erreur : \n\t%s", pingErr.Error())
	}

	log.Printf("Connection à la base de données réussie")
	return dbContext
}
