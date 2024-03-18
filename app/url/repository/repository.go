package repository

import "url-shorthand-with-golang/db/schema"

type UrlRepository interface {
	FindShortUrlRedis(shortUrl string) string
	FindShortUrlPostgresql(shortUrl string) string
	Create(url schema.Url) schema.Url
}
