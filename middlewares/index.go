package middlewares

import "github.com/gorilla/mux"

func RunMiddlewares(r *mux.Router) {
	r.Use(DefaultHeaders)
}
