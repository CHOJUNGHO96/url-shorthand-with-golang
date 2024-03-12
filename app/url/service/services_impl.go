package service

import (
	"github.com/go-playground/validator/v10"
	"time"
	"url-shorthand-with-golang/app/url/model"
	"url-shorthand-with-golang/app/url/repository"
	"url-shorthand-with-golang/db/schema"
)

type UrlServiceImpl struct {
	UrlRepository repository.UrlRepository
	Validate      *validator.Validate
}

func NewUrlServiceImpl(urlRepository repository.UrlRepository, validate *validator.Validate) UrlService {
	return &UrlServiceImpl{
		UrlRepository: urlRepository,
		Validate:      validate,
	}
}

func (s *UrlServiceImpl) FindShortUrl(shortUrl string) schema.Url {
	return s.UrlRepository.FindShortUrl(shortUrl)
}

func (s *UrlServiceImpl) Create(shorten model.Shorten) {
	err := s.Validate.Struct(shorten)
	if err != nil {
		panic(err)
	}
	urlSchema := schema.Url{UrlId: shorten.UrlId, LongUrl: shorten.LongUrl, ShortUrl: shorten.ShortUrl, IsEnable: 1, RegDate: time.Now()}
	s.UrlRepository.Create(urlSchema)
}
