package main

import (
	"github.com/NazarNintendo/cryptobot/logic"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	logic.Run()
}
