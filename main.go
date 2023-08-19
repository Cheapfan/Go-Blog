package main

import (
	"log"
	"os"

	"github.com/Cheapfan/Go-Blog/route"
	"github.com/Cheapfan/Go-Blog/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	route.Setup(app)
	app.Listen(":"+port)


}