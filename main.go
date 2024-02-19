package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func getArticle(w http.ResponseWriter, r *http.Request){
	// dateParam := chi.URLParam(r, "date")
	// slugParam := chi.URLParam(r, "slug")


	w.Write([]byte("data ini article"))
}

func main() {
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

		r.Get("/articles/{date}-{slug}", getArticle)
    
		http.ListenAndServe(":3000", r)
}