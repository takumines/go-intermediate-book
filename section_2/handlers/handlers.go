package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	models "github.com/takumines/go-intermediate-book/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

func HalloHandler(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Hello, World!\n")
	if err != nil {
		return
	}
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json \n", http.StatusBadRequest)
		return
	}

	article := reqArticle

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "internal server error \n", http.StatusInternalServerError)
		return
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	// コンパイルエラー回避用
	log.Println(page)

	articles := []models.Article{
		models.Article1,
		models.Article2,
	}

	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, "internal server error \n", http.StatusInternalServerError)
		return
	}
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(chi.URLParam(req, "articleID"))
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	}

	// コンパイルエラー回避用
	log.Println(articleId)

	article := models.Article1

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "internal server error \n", http.StatusInternalServerError)
		return
	}
}

func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json \n", http.StatusBadRequest)
		return
	}

	article := reqArticle
	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "internal server error \n", http.StatusInternalServerError)
		return
	}
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json \n", http.StatusBadRequest)
		return
	}

	comment := reqComment
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		http.Error(w, "internal server error \n", http.StatusInternalServerError)
		return
	}
}
