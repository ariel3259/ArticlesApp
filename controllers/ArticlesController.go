package controllers

import (
	"ArticlesApi/model"
	"ArticlesApi/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StartArticlesRouting(r *mux.Router) {
	//Init connection for Articles Services
	r.HandleFunc("/api/articles", getAll).Methods("GET")
	r.HandleFunc("/api/articles/{id}", getOne).Methods("GET")
	r.HandleFunc("/api/articles", save).Methods("POST")
	r.HandleFunc("/api/articles/{id}", modify).Methods("PUT")
	r.HandleFunc("/api/articles/{id}", delete).Methods("DELETE")
}

func getAll(w http.ResponseWriter, r *http.Request) {
	articles := services.GetArticles()
	json.NewEncoder(w).Encode(articles)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
	}
	article := services.GetOneArticle(uint(id))
	json.NewEncoder(w).Encode(article)
}

func save(w http.ResponseWriter, r *http.Request) {
	var article model.Articles
	json.NewDecoder(r.Body).Decode(&article)
	services.SaveArticle(&article)
	json.NewEncoder(w).Encode(article)
}

func modify(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var article model.Articles
	json.NewDecoder(r.Body).Decode(&article)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(nil)
	}
	services.ModifyArticle(uint(id), &article)
	json.NewEncoder(w).Encode(article)
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(nil)
	}
	services.DeleteArticle(uint(id))
	w.WriteHeader(204)
	json.NewEncoder(w).Encode(nil)
}
