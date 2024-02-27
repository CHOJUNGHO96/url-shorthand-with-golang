package shorturl

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GetShortUrl(c *gin.Context) {
	var url Url
	if c.ShouldBind(&url) == nil {
		log.Println(url.LongUrl)
	}

	c.String(200, "Success")
}
