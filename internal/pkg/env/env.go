package env

import (
	"github.com/joho/godotenv"
)

func LoadDefaultEnv() {
	_ = godotenv.Load(".env")
}
