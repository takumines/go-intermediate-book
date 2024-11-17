package services

import (
	"database/sql"
	"errors"
	"github.com/takumines/go-intermediate-book/apperrors"
	"github.com/takumines/go-intermediate-book/models"
	"github.com/takumines/go-intermediate-book/repositories"
)

// GetArticleService 記事詳細を取得するサービス
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {

	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = commentList

	return article, nil
}

// PostArticleService 記事を投稿するサービス
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert article")
		return models.Article{}, err
	}

	return newArticle, nil
}

// GetArticleListService 記事一覧を取得するサービス
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err = apperrors.NAData.Wrap(ErrNoData, "no data")
	}

	return articleList, nil
}

// PostNiceService いいねをするサービス
func (s *MyAppService) PostNiceService(articleID int) (models.Article, error) {
	err := repositories.UpdateNice(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")

			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}
