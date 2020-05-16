package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error during loading .env file")
	}
}

func GetVariableValue(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
