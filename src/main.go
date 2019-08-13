package main

import (
	"fmt"
	"log"
	"os"

	"./config"
	routerLoader "./router"

	"github.com/joho/godotenv"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	startApp()
}

func startApp() {
	router := config.SetupRouter()
	routerLoader.LoadRouter(router)
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)
	router.Run(serverString)
	fmt.Println("Server Up!!!")
}
