package schema

import "time"

type Url struct {
	UrlId    uint64    `json:"user_id" gorm:"primary_key"`
	ShortUrl string    `json:"short_url"`
	LongUrl  string    `json:"long_url"`
	IsEnable int       `json:"is_enable"`
	RegDate  time.Time `json:"reg_date"`
}

func (*Url) TableName() string {
	return "url"
}
