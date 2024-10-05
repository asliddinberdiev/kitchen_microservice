package main

import (
	"log"
	"net/http"

	"github.com/asliddinberdiev/kitchen_microservice/services/orders/handler"
	"github.com/asliddinberdiev/kitchen_microservice/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (svr *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("starting http server on: ", svr.addr)

	return http.ListenAndServe(svr.addr, router)
}
