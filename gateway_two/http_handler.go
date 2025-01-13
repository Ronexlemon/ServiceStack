package main

import (
	"context"
	"log"
	"net/http"

	common "github.com/ronexlemon/commons"
	pb "github.com/ronexlemon/otherservice/api"
)

type handlerGetMessageClient  struct{
	client pb.RouteGuideClient

}

func NewHandler(client pb.RouteGuideClient)*handlerGetMessageClient{
	return &handlerGetMessageClient{client}
}
func (h *handlerGetMessageClient) registerRoute(mux *http.ServeMux){
	mux.HandleFunc("POST /api/message/send",h.HandlerGetMessage)
}
func (h *handlerGetMessageClient) HandlerGetMessage(w http.ResponseWriter,r *http.Request){

	var mssgRequest *pb.MessageRequest
	log.Println(r.Body)
	if err := common.ReadJSON(r,&mssgRequest); err !=nil{
		common.WriteError(w,http.StatusBadRequest,err.Error())
	}
	res,err:= h.client.GetMessage(context.Background(),mssgRequest)
	if err != nil {
		common.WriteError(w,http.StatusInternalServerError,err.Error())
		return
	}
	common.WriteJSON(w,http.StatusOK,res)

}