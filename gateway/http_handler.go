package main

import (
	"net/http"
	//pb "github.com/RonexLemon/common/api"
)

type handler struct{

}
func NewHandler()*handler{
	return &handler{}
}
func (h *handler) registerRoute(mux *http.ServeMux){
	mux.HandleFunc("POST /api/customer/{customer}/orders",h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter ,r *http.Request){
	//log.Printf("Hllo")
	
	

}