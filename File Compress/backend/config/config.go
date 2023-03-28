package config

import (
	"github.com/joho/godotenv"
	"github.com/shuklaritvik06/GoProjects/filecompress/backend/database"
)

func Configure() {
	godotenv.Load(".env")
	database.DBConnect()
}
