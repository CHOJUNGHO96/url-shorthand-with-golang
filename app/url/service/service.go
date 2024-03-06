package service

import (
	"url-shorthand-with-golang/app/url/model"
)

type UrlService interface {
	Create(url model.Shorten)
}
