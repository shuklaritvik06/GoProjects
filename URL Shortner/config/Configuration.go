package config

import "github.com/joho/godotenv"

func Configure() {
	godotenv.Load(".env")
}
