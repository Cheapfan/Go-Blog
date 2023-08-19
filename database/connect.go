package database

import (
	"log"
	"os"

	"github.com/Cheapfan/Go-Blog/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	} else {
		log.Println("Connect successfully")
	}

	DB = database

	database.AutoMigrate( // To create the table inside the database
		&model.User{},
		&model.Blog{},
	)
}
