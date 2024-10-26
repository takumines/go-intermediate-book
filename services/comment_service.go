package services

import (
	"github.com/takumines/go-intermediate-book/config/database"
	"github.com/takumines/go-intermediate-book/models"
	"github.com/takumines/go-intermediate-book/repositories"
)

// PostCommentService コメントを投稿するサービス
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return models.Comment{}, err
	}

	comment, err = repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}
