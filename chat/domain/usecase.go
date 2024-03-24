package domain

import (
	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/sample_project/chat/client"
	"github.com/aliworkshop/sample_project/chat/client/data"
)

type RequestHandle func(c client.Client, request *data.Data)
type JoinHandler func(c client.Client, request *data.Data)

type ChatUc interface {
	Subscribe(userId uint64, ws gateway.WebSocketHandler) (client.Client, errors.ErrorModel)
	RegisterRequestHandler(t data.Type, handle RequestHandle)
	RegisterJoinHandler(t data.Type, handle JoinHandler)
	Start()
	GetClientByKey(key string) client.Client
}
