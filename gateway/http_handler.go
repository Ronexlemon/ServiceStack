package main

import (
	"net/http"

	pb "github.com/ronexlemon/commons/api"
	common "github.com/ronexlemon/commons"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}
func (h *handler) registerRoute(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customer/{customer}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	//log.Printf("Hllo")
	//pb.NewOrderServiceClient()
	
	customerid := r.PathValue("customerID")
	var items []*pb.ItemsWithQuantity
	if err:= common.ReadJSON(r,&items); err !=nil{
		common.WriteError(w,http.StatusBadRequest,err.Error())
	}
	h.client.CreateOrder(r.Context(),&pb.CreateOrderRequest{
		CustomerId: customerid,
		Items: items,
	})

}
