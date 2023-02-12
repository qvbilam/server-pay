package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pay/enum"
	"pay/global"
	"pay/model"
	"pay/utils"
	"time"
)

//type orderGoods struct {
//ID    int64
//Type  string
//Name  string
//Icon  string
//Price float64
//Count int64
//}

type OrderBusiness struct {
	ID         int64
	UserId     int64
	DeliveryId int64
	OrderSn    string
	TradeNo    string
	PayType    string
	ClientType string
	Status     string
	Amount     float64
	PayAmount  float64
	Remark     string
	PayResult  string
	PayTime    int64

	//// 商品信息
	//goods []*orderGoods
}

func (b *OrderBusiness) Create() (*model.Order, error) {
	tx := global.DB.Begin()

	type g struct {
		ID    int64
		Type  string
		Name  string
		Icon  string
		Price float64
		Count int64
	}
	gs := []g{
		{
			ID:    1,
			Type:  "exp",
			Name:  "测试经验",
			Icon:  "",
			Price: 10,
			Count: 1,
		},
		{
			ID:    2,
			Type:  "exp",
			Name:  "测试经验",
			Icon:  "",
			Price: 10,
			Count: 1,
		},
	}

	// 价格
	var amount float64
	for _, item := range gs {
		p := item.Price * float64(item.Count)
		amount += p
	}

	// 创建订单
	orderSn := utils.GenerateOrderSn(b.UserId)
	order := model.Order{
		UserModel: model.UserModel{
			UserID: b.UserId,
		},
		OrderSn:    orderSn,
		DeliveryId: b.DeliveryId,
		PayType:    b.PayType,
		ClientType: b.ClientType,
		Subject:    "购买商品",
		Status:     enum.PayStatusWait,
		Amount:     amount,
		Remark:     b.Remark,
	}
	if res := tx.Save(&order); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")
	}

	// 创建订单商品
	var goods []model.OrderGoods
	for _, item := range gs {
		goods = append(goods, model.OrderGoods{
			OrderId:   order.ID,
			GoodsId:   item.ID,
			GoodsType: item.Type,
			Name:      item.Name,
			Icon:      item.Icon,
			Price:     item.Price,
			Count:     item.Count,
		})
	}
	if res := tx.CreateInBatches(&goods, 1000); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")
	}
	tx.Commit()
	return &order, nil
}

func (b *OrderBusiness) Update() error {
	tx := global.DB.Begin()
	updataData := make(map[string]interface{})
	updataData["trade_no"] = b.TradeNo
	updataData["status"] = b.Status
	updataData["pay_result"] = b.PayResult
	updataData["pay_amount"] = b.PayAmount
	if b.PayTime != 0 {
		updataData["pay_time"] = time.Unix(b.PayTime, 0)
	}

	if res := tx.Model(&model.Order{}).Where(&model.Order{
		OrderSn: b.OrderSn,
	}).Updates(updataData); res.RowsAffected == 0 {
		tx.Rollback()
		return res.Error
	}

	tx.Commit()
	return nil
}
