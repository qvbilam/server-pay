package business

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"pay/enum"
	"pay/global"
	"pay/model"
	"pay/utils"
)

type OrderGoods struct {
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
	ClientType string
	Status     string
	Amount     float64
	PayAmount  *float64
	Remark     string
	PayResult  string
	PayTime    int64
	Ext        string
	Subject    string
	//// 商品信息
	Goods []*OrderGoods
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
		ClientType: b.ClientType,
		Subject:    b.Subject,
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
	for _, item := range b.Goods {
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

func (b *OrderBusiness) Update() (*model.Order, error) {
	tx := global.DB.Begin()
	entity := &model.Order{}
	if res := global.DB.Where(&model.Order{OrderSn: b.OrderSn}).First(&entity); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}
	if b.UserId != 0 && b.UserId != entity.UserID {
		tx.Rollback()
		return nil, status.Errorf(codes.Unauthenticated, "非法访问订单")
	}
	if b.PayType != "" {
		entity.PayType = b.PayType
	}
	if b.ClientType != "" {
		entity.ClientType = b.ClientType
	}
	if b.Status != "" {
		entity.Status = b.Status
	}
	if b.PayAmount != nil {
		entity.PayAmount = *b.PayAmount
	}
	if b.PayResult != "" {
		entity.PayResult = b.PayResult
	}
	if b.TradeNo != "" {
		entity.TradeNo = b.TradeNo
	}
	if b.PayTime != 0 {
		entity.PayTime = b.PayTime
	}

	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "订单变更失败")
	}
	tx.Commit()
	return entity, nil
}

func (b *OrderBusiness) sendOrderGoods(tx *gorm.DB) error {
	var OrderGoods []model.OrderGoods
	tx.Model(model.OrderGoods{}).Where(model.OrderGoods{OrderId: b.ID}).Find(&OrderGoods)

	for _, goods := range OrderGoods {
		fmt.Printf("%+v\n", goods)
		//global.UserServerClient.
		//global.UserServerClient.LevelExp()
	}

	fmt.Println(1)
	Transaction()
	fmt.Println(2)

	return nil
}
