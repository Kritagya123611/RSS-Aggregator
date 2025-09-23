package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Overload()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in environment variables")
	} else {
		fmt.Printf("Server will start at port: %s\n", port)
	}
}