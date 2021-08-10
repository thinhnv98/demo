package main

import (
	"demo/config"
	"demo/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("env/.dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Instance Server
	server := gin.Default()

	// Database
	database := config.Database{}
	db := database.InitDatabase()

	// Routes
	router := routes.Route{
		Db:     db,
		Server: *server,
	}
	router.Register()

	server.Run()
}
