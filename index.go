package main

import (
	"ArticlesApi/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	controllers.StartArticlesRouting(r)
	fmt.Println("Server online on port 8000")
	http.ListenAndServe(":8000", r)
}
