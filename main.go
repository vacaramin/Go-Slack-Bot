package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	LoadEnvVariables()
	//slacker.BotContext
	//bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"))
	fmt.Println(os.Getenv("SOCKET_TOKEN"))
}

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
