package env

import (
	"github.com/joho/godotenv"
)

func LoadDefaultEnv() {
	godotenv.Load(".env")
}
