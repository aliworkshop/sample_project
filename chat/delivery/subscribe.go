package delivery

import (
	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/sample_project/chat/client"
	"github.com/aliworkshop/sample_project/chat/domain"
	"github.com/gorilla/websocket"
)

type subscribeHandler struct {
	uc domain.ChatUc
}

func NewSubscribeHandler(useCase domain.ChatUc) gateway.Handler {
	handler := new(subscribeHandler)
	handler.uc = useCase
	return handler
}

func (h *subscribeHandler) Handle(request gateway.Requester) (any, errors.ErrorModel) {
	ws, err := request.Websocket()
	if err != nil {
		return nil, err
	}

	userId := request.GetCurrentAccountId()
	c, err := h.uc.Subscribe(userId, ws)
	if err != nil {
		return nil, err
	}

	c.Write(&client.WriteRequest{
		Type: websocket.TextMessage,
		Data: []byte("you've been subscribed successfully"),
	})

	return nil, nil
}
