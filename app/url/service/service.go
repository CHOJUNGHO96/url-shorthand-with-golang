package service

import (
	"url-shorthand-with-golang/app/url/model"
)

type UrlService interface {
	FindShortUrl(shortUrl string) string
	Create(url model.Shorten) string
}
