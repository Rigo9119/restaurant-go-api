package main

import (
	"log"
	"restaurant-go-api/cmd"
)

func main() {
	if err := cmd.StartServer(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
