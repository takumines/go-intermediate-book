package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/takumines/go-intermediate-book/config/database"
	"github.com/takumines/go-intermediate-book/controllers"
	"github.com/takumines/go-intermediate-book/services"

	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	s := services.NewMyAppService(db)
	c := controllers.NewMyAppController(s)

	r.Post("/article", c.PostArticleHandler)
	r.Get("/article/list", c.ArticleListHandler)
	r.Get("/article/{id:[0-9]+}", c.ArticleDetailHandler)
	r.Post("/article/nice", c.PostNiceHandler)
	r.Post("/comment", c.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
