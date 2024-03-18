package service

import (
	"url-shorthand-with-golang/app/url/model"
	"url-shorthand-with-golang/db/schema"
)

type UrlService interface {
	FindShortUrl(shortUrl string) schema.Url
	Create(url model.Shorten) string
}
