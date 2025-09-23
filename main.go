package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/cors" 
	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"http://*", "https://*"}, 
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:300, 
	}))

	router.Get(("/"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	router.Get(("/hello/{name}"), func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
	})

	http.ListenAndServe(":"+port, router)
}