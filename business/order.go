package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pay/enum"
	"pay/global"
	"pay/model"
	"pay/utils"
)

type orderGoods struct {
	ID    int64
	Type  string
	Name  string
	Icon  string
	Price float64
	Count int64
}

type OrderBusiness struct {
	ID         int64
	UserId     int64
	DeliveryId int64
	OrderSn    string
	TradeNo    string
	PayType    string
	Status     string
	Amount     float64
	Remark     string
	PayResult  string
	PayTime    int64

	// 商品信息
	goods []*orderGoods
}

func (b *OrderBusiness) Create() (*model.Order, error) {
	tx := global.DB.Begin()

	// 创建订单
	orderSn := utils.GenerateOrderSn(b.UserId)
	order := model.Order{
		UserModel: model.UserModel{
			UserID: b.UserId,
		},
		OrderSn:    orderSn,
		DeliveryId: b.DeliveryId,
		PayType:    b.PayType,
		Status:     enum.PayStatusWait,
		Amount:     b.Amount,
		Remark:     b.Remark,
	}
	if res := tx.Save(&order); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")
	}

	// 创建订单商品
	var goods []model.OrderGoods
	for _, item := range b.goods {
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

	return &order, nil
}
