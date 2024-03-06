package repository

import (
	"gorm.io/gorm"
	"url-shorthand-with-golang/db/schema"
)

type UrlRepositoryImpl struct {
	Db *gorm.DB
	//Gin *gin.Context
}

func NewUrlRepositoryImpl(Db *gorm.DB) UrlRepository {
	return &UrlRepositoryImpl{Db: Db}
}

func (r *UrlRepositoryImpl) Create(url schema.Url) {
	//db, _ := r.Gin.Get("db")
	result := r.Db.Create(&url)
	if result.Error != nil {
		panic(result.Error)
	}
}
