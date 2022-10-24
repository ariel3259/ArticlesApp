package main

import (
	"ArticlesApi/controllers"
	"ArticlesApi/database"
	"ArticlesApi/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	database.MakeModels(database.GetCon())
	middlewares.RunMiddlewares(r)
	controllers.StartRouters(r)
	handler := cors.AllowAll().Handler(r)
	fmt.Println("Server online on port 8000")

	http.ListenAndServe(":8000", handler)
}
