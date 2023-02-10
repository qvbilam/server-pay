package initialize

import (
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
	"pay/global"
	"time"
)

type dialConfig struct {
	host string
	port int64
	name string
}

type serverClientConfig struct {
	userDialConfig    *dialConfig
	contactDialConfig *dialConfig
	fileDialConfig    *dialConfig
}

func InitServer() {
	s := serverClientConfig{
		userDialConfig: &dialConfig{
			host: global.ServerConfig.UserServerConfig.Host,
			port: global.ServerConfig.UserServerConfig.Port,
			name: global.ServerConfig.UserServerConfig.Name,
		},
		contactDialConfig: &dialConfig{
			host: global.ServerConfig.ContactServerConfig.Host,
			port: global.ServerConfig.ContactServerConfig.Port,
			name: global.ServerConfig.ContactServerConfig.Name,
		},
	}

	s.initUserServer()
}

func clientOption() []retry.CallOption {
	opts := []retry.CallOption{
		retry.WithBackoff(retry.BackoffLinear(100 * time.Millisecond)), // 重试间隔
		retry.WithMax(3), // 最大重试次数
		retry.WithPerRetryTimeout(1 * time.Second),                                 // 请求超时时间
		retry.WithCodes(codes.NotFound, codes.DeadlineExceeded, codes.Unavailable), // 指定返回码重试
	}
	return opts
}

func (s *serverClientConfig) initUserServer() {
}
