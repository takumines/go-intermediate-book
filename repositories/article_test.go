package repositories_test

import (
	"github.com/takumines/go-intermediate-book/repositories"
	"testing"

	"github.com/takumines/go-intermediate-book/models"
)

func TestSelectArticleList(t *testing.T) {
	expectedNum := 2

	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != expectedNum {
		t.Errorf("got %d but want %d", len(got), expectedNum)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "first post",
				Contents: "This is my first blog",
				UserName: "takumi",
				Nice:     2,
			},
		},
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID:       2,
				Title:    "second post",
				Contents: "This is my second blog",
				UserName: "takumi",
				Nice:     3,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d", test.expected.ID, got.ID)
			}

			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s", test.expected.Title, got.Title)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s", test.expected.Contents, got.Contents)
			}

			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s", test.expected.UserName, got.UserName)
			}

			if got.Nice != test.expected.Nice {
				t.Errorf("Nice: get %d but want %d", test.expected.Nice, got.Nice)
			}
		})
	}
}

func TestInsertArticle(t *testing.T) {
	expectedArticle := models.Article{
		Title:    "third post",
		Contents: "This is my third blog",
		UserName: "takumi",
	}

	newArticle, err := repositories.InsertArticle(testDB, expectedArticle)
	if err != nil {
		t.Fatal(err)
	}

	if newArticle.Title != expectedArticle.Title {
		t.Errorf("Title: got %s but want %s", newArticle.Title, expectedArticle.Title)
	}

	if newArticle.Contents != expectedArticle.Contents {
		t.Errorf("Contents: got %s but want %s", newArticle.Contents, expectedArticle.Contents)
	}

	if newArticle.UserName != expectedArticle.UserName {
		t.Errorf("UserName: got %s but want %s", newArticle.UserName, expectedArticle.UserName)
	}

	// Clean up
	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?;
`
		testDB.Exec(sqlStr, expectedArticle.Title, expectedArticle.Contents, expectedArticle.UserName)
	})
}
