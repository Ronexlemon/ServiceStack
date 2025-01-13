package main

import (
	"errors"
	"net/http"

	common "github.com/ronexlemon/commons"
	pb "github.com/ronexlemon/commons/api"
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
	err := ValidateItems(items)
	if err !=nil{
		common.WriteError(w,http.StatusBadRequest,err.Error())
		return
	}
	o,err:=h.client.CreateOrder(r.Context(),&pb.CreateOrderRequest{
		CustomerId: customerid,
		Items: items,
	})
	if err !=nil{
		common.WriteError(w,http.StatusInternalServerError,err.Error())
	}
	common.WriteJSON(w,http.StatusOK,o)

}

func ValidateItems(items []*pb.ItemsWithQuantity)error{
	if len(items) == 0{
		return errors.New("Items must have atleast one item")
	}
	for _, item := range items {
		if item.Quantity <= 0 {
			return errors.New("Quantity must be greater than 0")
			}
		if item.ID == ""{
			return errors.New("Item ID is required")
		}
	}
	return nil
}
