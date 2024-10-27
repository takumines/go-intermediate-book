package api

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/takumines/go-intermediate-book/controllers"
	"github.com/takumines/go-intermediate-book/services"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	s := services.NewMyAppService(db)
	ac := controllers.NewArticleController(s)
	cc := controllers.NewCommentController(s)

	r.Post("/article", ac.PostArticleHandler)
	r.Get("/article/list", ac.ArticleListHandler)
	r.Get("/article/{id:[0-9]+}", ac.ArticleDetailHandler)
	r.Post("/article/nice", ac.PostNiceHandler)
	r.Post("/comment", cc.PostCommentHandler)

	return r
}
