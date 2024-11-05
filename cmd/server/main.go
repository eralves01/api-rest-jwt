package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eralves01/api-rest-jwt/configs"
)

func main() {
	fileConfigName := os.Getenv("APP_ENV")
	if fileConfigName == "" {
		fileConfigName = "default"
	}

	if err := configs.LoadConfig(fileConfigName); err != nil {
		log.Fatal("Error loading config: ", err)
	}

	fmt.Println(configs.AppConfig)
	fmt.Println("Server running on port:", configs.AppConfig.Server.Port)
	fmt.Println("Server running on port:", configs.AppConfig.Database.Host)
}
