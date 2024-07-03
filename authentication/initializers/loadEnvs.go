package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
}
