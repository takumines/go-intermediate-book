package services

import "github.com/takumines/go-intermediate-book/models"

type ArticleServicer interface {
	GetArticleService(articleID int) (models.Article, error)
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostNiceService(articleID int) (models.Article, error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
