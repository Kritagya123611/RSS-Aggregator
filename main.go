package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/Kritagya123611/Go-Backened/internal/auth"
	"github.com/Kritagya123611/Go-Backened/internal/products"
)

func main() {
	err := godotenv.Overload()
	if err != nil {
		log.Printf("No .env file found, using defaults")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" 
	}

	fmt.Printf("Server will start at port: %s\n", port)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Post("/signup", auth.SignupHandler)
	router.Post("/login",auth.LoginHandler)
	router.Get("/products", products.AllProducts)
	router.Get("/products/{id}", products.GetProductByID) 
	http.ListenAndServe(":"+port, router)
}
