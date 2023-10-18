package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	proto "pay/api/qvbilam/pay/v1"
	"pay/business"
)

type PayServer struct {
	proto.UnimplementedPayServer
}

func RequestGoodsConvert(goods []*proto.CreateGoodsRequest) []*business.OrderGoods {
	if goods == nil {
		return nil
	}
	var res []*business.OrderGoods
	for _, g := range goods {
		res = append(res, &business.OrderGoods{
			ID:    g.Id,
			Type:  g.Type,
			Name:  g.Name,
			Icon:  g.Icon,
			Price: float64(g.Price),
			Count: g.Count,
		})
	}
	return res

}

func (s *PayServer) CreateOrder(ctx context.Context, request *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	b := business.OrderBusiness{
		UserId:     request.UserId,
		DeliveryId: request.DeliveryId,
		PayType:    request.PayType,
		ClientType: request.ClientType,
		Amount:     float64(request.Amount),
		Remark:     request.Remark,
		Ext:        request.Ext,
		Subject:    request.Subject,
		Goods:      RequestGoodsConvert(request.Goods),
	}
	order, err := b.Create()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "创建支付订单失败")
	}
	return &proto.OrderResponse{
		UserID:  order.UserID,
		OrderSn: order.OrderSn,
		PayType: order.PayType,
		Amount:  float32(order.Amount),
		Subject: order.Subject,
		Remark:  order.Remark,
	}, nil
}

func (s *PayServer) UpdateOrder(ctx context.Context, request *proto.UpdateOrderRequest) (*proto.OrderResponse, error) {
	payAmount := float64(request.PayAmount)
	b := business.OrderBusiness{
		OrderSn:   request.OrderSn,
		TradeNo:   request.TradeNo,
		Status:    request.Status,
		PayResult: request.PayResult,
		PayAmount: &payAmount,
		PayTime:   request.PayTime,
	}

	if _, err := b.Update(); err != nil {
		return nil, err
	}

	return &proto.OrderResponse{}, nil
}

func (s *PayServer) ApplyOrder(ctx context.Context, request *proto.ApplyOrderRequest) (*proto.OrderResponse, error) {
	b := business.OrderBusiness{
		OrderSn:    request.OrderSn,
		PayType:    request.PayType,
		ClientType: request.ClientType,
		UserId:     request.UserID,
	}
	entity, err := b.Update()
	if err != nil {
		return nil, err
	}
	return &proto.OrderResponse{
		UserID:     entity.UserID,
		OrderSn:    entity.OrderSn,
		PayType:    entity.PayType,
		ClientType: entity.ClientType,
		Amount:     float32(entity.Amount),
		Subject:    entity.Subject,
		Remark:     entity.Remark,
	}, nil
}
