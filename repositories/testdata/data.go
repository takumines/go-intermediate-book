package testdata

import "github.com/takumines/go-intermediate-book/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "first post",
		Contents: "This is my first blog",
		UserName: "takumi",
		Nice:     2,
	},
	models.Article{
		ID:       2,
		Title:    "second post",
		Contents: "This is my second blog",
		UserName: "takumi",
		Nice:     3,
	},
}
