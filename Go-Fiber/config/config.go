package config

import (
	"github.com/joho/godotenv"
	"github.com/shuklaritvik06/GoProjects/fiber/database"
)

func Configure() {
	godotenv.Load(".env")
	database.Connect()
}
