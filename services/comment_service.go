package services

import (
	"github.com/takumines/go-intermediate-book/apperrors"
	"github.com/takumines/go-intermediate-book/models"
	"github.com/takumines/go-intermediate-book/repositories"
)

// PostCommentService コメントを投稿するサービス
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
