package service

import (
	"fmt"
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

func (s *UrlServiceImpl) Create(shorten model.Shorten) string {
	err := s.Validate.Struct(shorten)
	if err != nil {
		panic(err)
	}
	urlSchema := schema.Url{UrlId: shorten.UrlId, LongUrl: shorten.LongUrl, ShortUrl: shorten.ShortUrl, IsEnable: 1, RegDate: time.Now()}
	urlInfo := s.UrlRepository.Create(urlSchema)
	return fmt.Sprintf("url : http://127.0.0.1:8080/%s", urlInfo.ShortUrl) // 실제 운영시에는 단축기 프로젝트 도메인으로 설정
}
