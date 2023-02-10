package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"pay/config"
)

var (
	DB           *gorm.DB
	Redis        redis.Client
	ES           *elastic.Client
	ServerConfig *config.ServerConfig
)
