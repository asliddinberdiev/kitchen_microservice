package handler

import (
	"net/http"

	"github.com/asliddinberdiev/kitchen_microservice/services/common/genproto/orders"
	"github.com/asliddinberdiev/kitchen_microservice/services/orders/types"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{orderService: orderService}
}


func (h *OrdersHttpHandler) RegisterRouter (router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	if err := h.orderService.CreateOrder(r.Context(), order); err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}