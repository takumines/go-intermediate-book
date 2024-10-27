package main

import (
	"github.com/takumines/go-intermediate-book/config/database"
	"github.com/takumines/go-intermediate-book/controllers"
	"github.com/takumines/go-intermediate-book/routers"
	"github.com/takumines/go-intermediate-book/services"

	"log"
	"net/http"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	s := services.NewMyAppService(db)
	c := controllers.NewMyAppController(s)
	r := routers.NewRouter(c)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
