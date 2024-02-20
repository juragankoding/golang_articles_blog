package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"example.com/chin/connection"
	"example.com/chin/database"
	"example.com/chin/env"
	"github.com/go-playground/validator"
)

type requestArticle struct {
	Name        string `validate:"required", json:"name"`
	Description string `validate:"required", json:"description"`
}

type response struct {
	data any
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
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

	connection, err := connection.Connect(env.GetMongoDBURL())

	adapter := database.ArticleMongoDb{Client: connection}

	resultAdapter := adapter.InsertArticleIntoDatabase(
		database.Article{
			Name:        article.Name,
			Description: article.Description,
		})

	data, err := json.Marshal(response{
		data: resultAdapter,
	})

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	r.URL.Query()

	// dateParam := chi.URLParam(r, "date")
	// slugParam := chi.URLParam(r, "slug")

	w.Write([]byte("data ini article"))
}
