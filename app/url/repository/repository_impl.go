package repository

import (
	"gorm.io/gorm"
	"url-shorthand-with-golang/db/schema"
)

type UrlRepositoryImpl struct {
	Db *gorm.DB
}

func NewUrlRepositoryImpl(Db *gorm.DB) UrlRepository {
	return &UrlRepositoryImpl{Db: Db}
}

func (r *UrlRepositoryImpl) FindShortUrl(shortUrl string) schema.Url {
	var urlInfo schema.Url
	if err := r.Db.Where("short_url = ?", shortUrl).First(&urlInfo).Error; err != nil {
		panic(err)
	} else {
		return urlInfo
	}

}

func (r *UrlRepositoryImpl) Create(url schema.Url) {
	result := r.Db.Create(&url)
	if result.Error != nil {
		panic(result.Error)
	}
}
