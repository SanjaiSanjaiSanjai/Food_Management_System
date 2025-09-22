package db

import (
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	"Food_Delivery_Management/utils"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConfig() {
	err := godotenv.Load()
	// err is not nil return error
	utils.IsNotNilError(err, "DB config", "godotenv error")

	// connection variables pg to go
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DBNAME")

	// DB connecting string
	dsn := fmt.Sprintf(
		"host=%s user=%s  password=%s dbname=%s port=%s  sslmode=disable",
		host, username, password, database, port)

	// connected pg to gorm
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// error is not nil return error
	utils.IsNotNilError(error, "DB config", "database connection error")
	customlogger.Log.Info("[DB config]: Database successfully connected")
	DB = db
}
