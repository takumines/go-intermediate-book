package main

import (
	"github.com/takumines/go-intermediate-book/api"
	"github.com/takumines/go-intermediate-book/config/database"
	"log"
	"net/http"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
