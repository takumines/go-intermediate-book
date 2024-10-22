package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/takumines/go-intermediate-book/handlers"

	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/hello", handlers.HelloHandler)

	r.Post("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.ArticleListHandler)
	r.Get("/article/{id:[0-9]+}", handlers.ArticleDetailHandler)
	r.Post("/article/nice", handlers.PostNiceHandler)
	r.Post("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
