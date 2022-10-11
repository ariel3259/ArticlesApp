package services

import (
	"ArticlesApi/database"
	"ArticlesApi/model"

	"gorm.io/gorm"
)

var db *gorm.DB

func SetConForArticles(_db *gorm.DB) {
	db = _db
	database.MakeModels(db)
}

func GetArticles() []model.Articles {
	var articles []model.Articles
	db.Find(&articles, model.Articles{State: true})
	return articles
}

func GetOneArticle(id uint) model.Articles {
	var article model.Articles
	db.First(&article, model.Articles{State: true, Model: gorm.Model{ID: id}})
	return article
}

func SaveArticle(article *model.Articles) {
	article.State = true
	db.Save(article)
}

func ModifyArticle(id uint, article *model.Articles) {
	article.ID = id
	db.Save(article)
}

func DeleteArticle(id uint) {
	var article model.Articles
	db.First(&article, model.Articles{Model: gorm.Model{ID: id}})
	article.State = false
	db.Save(&article)
}
