package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
	url "url-shorthand-with-golang/app/url"
	"url-shorthand-with-golang/app/url/repository"
	"url-shorthand-with-golang/app/url/service"
	core "url-shorthand-with-golang/core"
	db "url-shorthand-with-golang/db"
	"url-shorthand-with-golang/db/schema"
)

func main() {
	// 강제적으로 로그의 색상을 지정
	gin.ForceConsoleColor()

	// 유효성 검사기 생성
	validate := validator.New()

	// 라우터 생성
	r := gin.Default()

	// 라우터 그룹 생성
	config := core.SetConfig()

	// Postgresql 연결
	postgres := db.Postgresql(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul", config.PostgresServer, config.PostgresUser, config.PostgresPassword, config.PostgresDb, config.PostgresPort))

	// 데이터베이스 마이그레이션
	postgres.AutoMigrate(&schema.Url{})

	// LoggerWithFormatter 미들웨어는 gin.DefaultWriter에 로그를 작성
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

	urlRepository := repository.NewUrlRepositoryImpl(postgres)
	urlService := service.NewUrlServiceImpl(urlRepository, validate)
	urlEndpoint := url.NewUrlEndpoint(urlService)

	// 라우트 그룹화
	v1 := r.Group("/api/v1")
	{
		v1.POST("/url/shorten", urlEndpoint.PostShorten)
	}
	r.Run(":8080")

}
