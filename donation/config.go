package donation

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpnPublicKey string
	OpnSecretKey string
}

func GetEnv() (config Config) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	opnPublicKey := os.Getenv("PUBLIC_KEY_OPN")
	ppnSecretKey := os.Getenv("SECRET_KEY_OPN")

	return Config{OpnPublicKey: ppnSecretKey, OpnSecretKey: opnPublicKey}
}