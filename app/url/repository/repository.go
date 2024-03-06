package repository

import "url-shorthand-with-golang/db/schema"

type UrlRepository interface {
	Create(url schema.Url)
}
