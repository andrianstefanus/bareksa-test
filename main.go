package main

import (
	"bareksa-test/database"
	d "bareksa-test/delivery"
	r "bareksa-test/repository"
	u "bareksa-test/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	// news
	newsRepository := r.InitiateNewsRepository(database.Databases)
	newsUsecase := u.InitiateNewsUsecase(newsRepository)
	d.InitiateNewsHandler(route, newsUsecase)

	// tags
	tagsRepository := r.InitiateTagsRepository(database.Databases)
	tagsUsecase := u.InitiateTagsUsecase(tagsRepository)
	d.InitiateTagsHandler(route, tagsUsecase)

	// add route
	http.Handle("/", route)

	// Run
	log.Println("Server is running:", 5074)
	log.Fatal(http.ListenAndServe(":5074", nil))
}
