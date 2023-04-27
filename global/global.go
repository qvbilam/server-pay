package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	userProto "pay/api/qvbilam/user/v1"
	"pay/config"
)

var (
	DB               *gorm.DB
	Redis            redis.Client
	ES               *elastic.Client
	ServerConfig     *config.ServerConfig
	UserServerClient userProto.UserClient
)
