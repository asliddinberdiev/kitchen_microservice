package types

import (
	"context"

	"github.com/asliddinberdiev/kitchen_microservice/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *orders.Order) error
}
