package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"url-shorthand-with-golang/app/shorturl"
)

func main() {
	// 강제적으로 로그의 색상을 지정
	gin.ForceConsoleColor()

	r := gin.Default()

	// LoggerWithFormatter 미들웨어는 gin.DefaultWriter에 로그를 작성합니다.
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 사용자 정의 형식
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// 라우트 그룹화
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ShortUrl", shorturl.GetShortUrl)
	}
	r.Run(":8080")

}
