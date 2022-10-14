package controllers

import "github.com/gorilla/mux"

func StartRouters(r *mux.Router) {
	StartArticlesRouting(r)
}
