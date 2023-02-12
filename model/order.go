package model

import "time"

type Order struct {
	IDModel
	UserModel
	OrderSn    string `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:订单编号;index"`
	TradeNo    string `gorm:"type:varchar(100) NOT NULL DEFAULT '';comment:交易号;index"`
	DeliveryId int64  `gorm:"type:int NOT NULL DEFAULT 0 comment '收货id';index"`
	//Delivery   *Delivery
	PayType    string     `gorm:"type:varchar(20) NOT NULL DEFAULT 'alipay';comment:支付方式:alipay,wechat"`
	ClientType string     `gorm:"type:varchar(20) NOT NULL DEFAULT 'alipay';comment:客户端类型:web,Android,iOS"`
	Status     string     `gorm:"type:varchar(20) NOT NULL DEFAULT 'WAIT';comment:WAIT(等待支付),SUCCESS(成功),CLOSED(关闭),FINISHED(交易结束)"`
	Amount     float64    `gorm:"NOT NULL DEFAULT 0;comment:总金额"`
	PayAmount  float64    `gorm:"NOT NULL DEFAULT 0;comment:支付金额"`
	Subject    string     `gorm:"type:varchar(20) NOT NULL DEFAULT '';comment:主题"`
	Remark     string     `gorm:"type:varchar(20) NOT NULL DEFAULT '';comment:备注"`
	PayResult  string     `gorm:"type:varchar(2048) NOT NULL DEFAULT '';comment:支付结果"`
	PayTime    *time.Time `gorm:"type:datetime;comment:支付时间"`
	DateModel
}

// OrderGoods 订单商品关联
type OrderGoods struct {
	IDModel
	OrderId   int64  `gorm:"type:int NOT NULL DEFAULT 0;comment:订单id;index"`
	GoodsType string `gorm:"type:varchar(100) NOT NULL DEFAULT '';comment:商品类型;index:goods_type_id"`
	GoodsId   int64  `gorm:"type:int NOT NULL DEFAULT 0;comment:商品id;index:goods_type_id"`

	Name  string  `gorm:"type:varchar(100) NOT NULL DEFAULT '';comment:商品名称"`
	Icon  string  `gorm:"type:varchar(200) NOT NULL DEFAULT '';comment:图标"`
	Price float64 `gorm:"NOT NULL DEFAULT 0;comment:商品价格"`
	Count int64   `gorm:"type:int NOT NULL DEFAULT 0;comment:商品数量"`
	DateModel
}

// Delivery 收货地址
type Delivery struct {
	IDModel
	UserModel
	Name      string `gorm:"type:varchar(20) NOT NULL DEFAULT '';comment:收件人昵称"`
	Mobile    string `gorm:"type:varchar(11) NOT NULL DEFAULT '';comment:手机号"`
	Address   string `gorm:"type:varchar(100) NOT NULL DEFAULT '';comment:收货地址"`
	Tag       string `gorm:"type:varchar(20) NOT NULL DEFAULT '';comment:标签"`
	IsDefault bool   `gorm:"NOT NULL DEFAULT 0;comment:是否默认"`
	DateModel
}
