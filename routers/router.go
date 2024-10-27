package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/takumines/go-intermediate-book/controllers"
)

func NewRouter(c *controllers.MyAppController) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/article", c.PostArticleHandler)
	r.Get("/article/list", c.ArticleListHandler)
	r.Get("/article/{id:[0-9]+}", c.ArticleDetailHandler)
	r.Post("/article/nice", c.PostNiceHandler)
	r.Post("/comment", c.PostCommentHandler)

	return r
}
