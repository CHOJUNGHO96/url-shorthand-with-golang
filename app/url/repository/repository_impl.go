package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
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
	bytes, _ := json.Marshal(url)
	err := r.Redis.Set(ctx, strconv.FormatUint(url.UrlId, 10), string(bytes), time.Hour*24).Err()
	if err != nil {
		panic(err)
	}
}
