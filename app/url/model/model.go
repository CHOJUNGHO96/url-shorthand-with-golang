package model

type ShortUrl struct {
	ShortUrl string `form:"ShortUrl"`
}

type Shorten struct {
	LongUrl  string `form:"LongUrl"`
	UrlId    uint64
	ShortUrl string
}
