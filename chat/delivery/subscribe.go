package delivery

import (
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/sample_project/chat/client"
	"github.com/aliworkshop/sample_project/chat/domain"
	"github.com/gorilla/websocket"
)

type subscribeHandler struct {
	handlerlib.HandlerModel
	uc domain.ChatUc
}

func NewSubscribeHandler(handlerModel handlerlib.HandlerModel, useCase domain.ChatUc) handlerlib.HandlerModel {
	handler := new(subscribeHandler)
	handler.HandlerModel = handlerModel
	handler.uc = useCase
	handler.SetHandlerFunc(handler.handle)
	return handler
}

func (h *subscribeHandler) handle(request handlerlib.RequestModel, args ...interface{}) (interface{}, errorslib.ErrorModel) {
	ws, err := h.Upgrade(request)
	if err != nil {
		return nil, errorslib.Internal(err)
	}

	userId := request.GetAuth().GetClaim().GetUserId()
	c, err := h.uc.Subscribe(userId, ws)
	if err != nil {
		return nil, errorslib.Internal(err)
	}

	c.Write(&client.WriteRequest{
		Type: websocket.TextMessage,
		Data: []byte("you've been subscribed successfully"),
	})

	h.Respond(request, handlerlib.StatusOK, nil)
	request.SetResponded(true)
	return nil, nil
}
