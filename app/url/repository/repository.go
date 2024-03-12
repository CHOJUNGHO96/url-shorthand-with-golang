package repository

import "url-shorthand-with-golang/db/schema"

type UrlRepository interface {
	FindShortUrl(shortUrl string) schema.Url
	Create(url schema.Url)
}
