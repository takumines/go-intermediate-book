package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
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
	_, err := io.WriteString(w, "Post Article\n")
	if err != nil {
		return
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("http request: %v\n", req)
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

	resStr := fmt.Sprintf("Article List (page %d)\n", page)
	_, err := io.WriteString(w, resStr)
	if err != nil {
		return
	}
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("http request: %v\n", req.PathValue("articleID"))
	articleId, err := strconv.Atoi(chi.URLParam(req, "articleID"))
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	}
	resStr := fmt.Sprintf("Article No.%d\n", articleId)
	_, err = io.WriteString(w, resStr)
	if err != nil {
		return
	}
}

func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Posting Nice Article\n")
	if err != nil {
		return
	}
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Comment\n")
	if err != nil {
		return
	}
}
