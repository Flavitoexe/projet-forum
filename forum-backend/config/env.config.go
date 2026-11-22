package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envPath := findEnvFile()
	if envPath == "" {
		log.Println("Aucun fichier .env trouvé.\n\tUtilisation des variables d'environnement système.")
		return
	}

	errLoad := godotenv.Load(envPath)
	if errLoad != nil {
		log.Println("Impossible de charger le fichier .env.\n\tUtilisation des variables d'environnement système.")
	}
}

func findEnvFile() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	for dir := cwd; ; dir = filepath.Dir(dir) {
		candidate := filepath.Join(dir, ".env")
		if info, statErr := os.Stat(candidate); statErr == nil && !info.IsDir() {
			return candidate
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
	}
}

func GetEnvWithDefault(key, defaultValue string) string {
	envVar, envErr := os.LookupEnv(key)
	if !envErr {
		return defaultValue
	}
	return envVar
}

func GetRequiredEnv(key string) string {
	envVar, envErr := os.LookupEnv(key)
	if !envErr || envVar == "" {
		log.Fatalf("Erreur configuration - Variable d'environnement manquante : %s", key)
	}
	return envVar
}
