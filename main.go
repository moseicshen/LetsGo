package main

import (
	"LetsGo/config"
	"LetsGo/models"
	"LetsGo/routers"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port, exist := os.LookupEnv("PORT")
	if !exist {
		panic("Missing environment settings")
	}
	// initialize database
	config.InitDB()
	// update schema
	err = config.DB.AutoMigrate(&models.List{})
	if err != nil {
		panic("failed to migrate schema")
	}
	// run server
	engine := routers.InitRouter()
	portStr := fmt.Sprintf(":%s", port)
	err = engine.Run(portStr)
	if err != nil {
		return
	}
}
