package core

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Config 구조체 정의
type Config struct {
	PostgresServer   string `mapstructure:"POSTGRES_SERVER"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	RedisPort        string `mapstructure:"REDIS_PORT"`
	RedisServer      string `mapstructure:"REDIS_SERVER"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	RedisExpireTime  int    `mapstructure:"REDIS_EXPIRE_TIME"`
}

func SetConfig() *Config {
	c := Config{}
	// .env 파일 로드
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
	}

	// 환경 변수 로드
	viper.AutomaticEnv()

	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &c
}
