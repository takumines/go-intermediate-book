package repositories_test

import (
	"github.com/takumines/go-intermediate-book/models"
	"github.com/takumines/go-intermediate-book/repositories"
	"testing"
)

func TestSelectCommentList(t *testing.T) {
	articleId := 1
	expectedCommentNum := 1

	got, err := repositories.SelectCommentList(testDB, articleId)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != expectedCommentNum {
		t.Errorf("got %d but want %d", len(got), expectedCommentNum)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "This is my first comment",
	}

	got, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}

	if got.ArticleID != comment.ArticleID {
		t.Errorf("ArticleID: got %d but want %d", got.ArticleID, comment.ArticleID)
	}

	if got.Message != comment.Message {
		t.Errorf("Message: got %s but want %s", got.Message, comment.Message)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where comment_id = ?;
		`
		testDB.Exec(sqlStr, got.CommentID)
	})
}
