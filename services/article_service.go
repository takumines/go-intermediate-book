package services

import (
	"github.com/takumines/go-intermediate-book/config/database"
	"github.com/takumines/go-intermediate-book/models"
	"github.com/takumines/go-intermediate-book/repositories"
)

// GetArticleService 記事詳細を取得するサービス
func GetArticleService(articleID int) (models.Article, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return models.Article{}, err
	}

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = commentList

	return article, nil
}

// PostArticleService 記事を投稿するサービス
func PostArticleService(article models.Article) (models.Article, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return models.Article{}, err
	}

	article, err = repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

// GetArticleListService 記事一覧を取得するサービス
func GetArticleListService(page int) ([]models.Article, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

// PostNiceService いいねをするサービス
func PostNiceService(articleID int) (models.Article, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return models.Article{}, err
	}

	err = repositories.UpdateNice(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}
