package services

import (
	"ArticlesApi/database"
	"ArticlesApi/model"

	"gorm.io/gorm"
)

var db *gorm.DB = database.GetCon()

func GetArticles(offset int, limit int, total *int64) []model.Articles {
	var articles []model.Articles
	db.Model(model.Articles{}).Count(total)
	db.Offset(offset).Limit(limit).Find(&articles, model.Articles{State: true})
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
