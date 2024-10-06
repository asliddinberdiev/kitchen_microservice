package handler

import (
	"context"

	"github.com/asliddinberdiev/kitchen_microservice/services/common/genproto/orders"
	"github.com/asliddinberdiev/kitchen_microservice/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 42,
		ProductID:  42,
		Quantity:   10,
	}

	if err := h.ordersService.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	orderList := h.ordersService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: orderList,
	}

	return res, nil
}
