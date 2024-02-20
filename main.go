package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/chin/endpoint"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/articles/{date}-{slug}", endpoint.GetArticle)
	r.Post("/article", endpoint.CreateArticle)

	fmt.Println("Server berlajan di port :3000")

	err = http.ListenAndServe(":3000", r)

	if err != nil {
		panic(err)
	}

}
