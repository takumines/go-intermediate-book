package main

import (
	"github.com/takumines/go-intermediate-book/section_2/handlers"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/hello", handlers.HalloHandler)
	r.Post("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.ArticleListHandler)
	r.Get("/article/{articleID:[0-9]+}", handlers.ArticleDetailHandler)
	r.Post("/article/nice", handlers.PostNiceArticleHandler)
	r.Post("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
