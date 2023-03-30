package config

import (
	"github.com/joho/godotenv"
	"github.com/shuklaritvik06/GoProjects/gochi/database"
)

func Configure() {
	godotenv.Load(".env")
	database.Connect()
}
