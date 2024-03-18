package url

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shorthand-with-golang/app/url/model"
	"url-shorthand-with-golang/app/url/service"
)

type Endpoint struct {
	urlService service.UrlService
}

func NewUrlEndpoint(urlService service.UrlService) Endpoint {
	return Endpoint{urlService: urlService}
}

func (e *Endpoint) GetShorUrl(c *gin.Context) {
	shortUrl := c.Params.ByName("short_url")
	userInfo := e.urlService.FindShortUrl(shortUrl)
	c.Redirect(http.StatusMovedPermanently, userInfo.LongUrl)
}

func (e *Endpoint) PostShorten(c *gin.Context) {
	var url model.Shorten
	if c.ShouldBind(&url) == nil {
		hashData := GetHashUrl(url.LongUrl)[0:7]      // ef7efc9
		base62DecodeData := GetDecodeBase62(hashData) // 2309683996913
		url.ShortUrl = hashData
		url.UrlId = uint64(base62DecodeData)
		c.JSON(http.StatusOK, e.urlService.Create(url))
	}
}
