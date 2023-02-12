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

func (s *PayServer) CreateOrder(ctx context.Context, request *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	b := business.OrderBusiness{
		UserId:     request.UserId,
		DeliveryId: request.DeliveryId,
		PayType:    request.PayType,
		ClientType: request.ClientType,
		Remark:     request.Remark,
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
	b := business.OrderBusiness{
		OrderSn:   request.OrderSn,
		TradeNo:   request.TradeNo,
		Status:    request.Status,
		PayResult: request.PayResult,
		PayAmount: float64(request.PayAmount),
		PayTime:   request.PayTime,
	}

	if err := b.Update(); err != nil {
		return nil, err
	}

	return &proto.OrderResponse{}, nil
}
