package env

import (
	"github.com/joho/godotenv"
)

func LoadDefaultEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
