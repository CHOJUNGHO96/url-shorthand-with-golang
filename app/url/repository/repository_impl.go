package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
	"url-shorthand-with-golang/db/schema"
)

var ctx = context.Background()

type UrlRepositoryImpl struct {
	Db    *gorm.DB
	Redis *redis.Client
}

func NewUrlRepositoryImpl(Db *gorm.DB, Redis *redis.Client) UrlRepository {
	return &UrlRepositoryImpl{Db: Db, Redis: Redis}
}

func (r *UrlRepositoryImpl) FindShortUrlRedis(shortUrl string) string {
	return r.Redis.Get(ctx, shortUrl).Val()
}

func (r *UrlRepositoryImpl) FindShortUrlPostgresql(shortUrl string) string {
	var urlInfo schema.Url
	if err := r.Db.Where("short_url = ?", shortUrl).First(&urlInfo).Error; err != nil {
		panic(err)
	} else {
		bytes, _ := json.Marshal(urlInfo)
		err := r.Redis.Set(ctx, urlInfo.ShortUrl, string(bytes), time.Hour*24).Err()
		if err != nil {
			panic(err)
		}
		return urlInfo.LongUrl
	}
}

func (r *UrlRepositoryImpl) Create(url schema.Url) schema.Url {
	result := r.Db.Create(&url)
	if result.Error != nil {
		panic(result.Error)
	}
	bytes, _ := json.Marshal(url)
	err := r.Redis.Set(ctx, url.ShortUrl, string(bytes), time.Hour*24).Err()
	if err != nil {
		panic(err)
	}
	return url
}
