package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/asliddinberdiev/kitchen_microservice/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (svr *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  3123,
			Quantity:   2,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		orderList, err := c.GetOrders(ctx, &orders.GetOrderRequest{CustomerID: 42})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		temp := template.Must(template.New("orders").Parse(ordersTemplate))
		if err := temp.Execute(w, orderList.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}
	})

	log.Println("starting http server on: ", svr.addr)
	return http.ListenAndServe(svr.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
