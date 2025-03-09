package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := fmt.Sprintln(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil{
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connected successfully!")
}