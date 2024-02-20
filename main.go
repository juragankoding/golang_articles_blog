package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

type requestArticle struct {
	Name        string `validate:"required", json:"name"`
	Description string `validate:"required", json:"description"`
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	// dateParam := chi.URLParam(r, "date")
	// slugParam := chi.URLParam(r, "slug")

	w.Write([]byte("data ini article"))
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	bodyData, _ := ioutil.ReadAll(r.Body)

	var article requestArticle

	json.Unmarshal(bodyData, &article)

	fmt.Println(article)

	validate := validator.New()

	err := validate.Struct(article)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	data, err := json.Marshal(article)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/articles/{date}-{slug}", getArticle)
	r.Post("/article", createArticle)

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server berlajan di port :3000")
}
