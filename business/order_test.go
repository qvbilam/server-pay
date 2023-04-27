package business

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	userProto "pay/api/qvbilam/user/v1"
	"pay/global"
	"testing"
)

func initUserClient() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", "127.0.0.1", 9801),
		grpc.WithInsecure())
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "grpc-user-server", err)
	}
	//
	userClient := userProto.NewUserClient(conn)
	global.UserServerClient = userClient
}

func initDBClient() {
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	database := "qvbilam_pay"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不带表名
		},
	})
	if err != nil {
		panic(any(err))
	}
	global.DB = db
}

func TestOrderBusiness_Update(t *testing.T) {
	initDBClient()
	initUserClient()

	tx := global.DB
	b := OrderBusiness{
		ID: 1,
	}
	b.sendOrderGoods(tx)
}
