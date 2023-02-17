package domain

import (
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/sample_project/chat/client"
	"github.com/aliworkshop/sample_project/chat/client/data"
)

type RequestHandle func(c client.Client, request *data.Data)
type JoinHandler func(c client.Client, request *data.Data)

type ChatUc interface {
	Subscribe(userId uint64, ws handlerlib.WebSocketModel) (client.Client, errorslib.ErrorModel)
	RegisterRequestHandler(t data.Type, handle RequestHandle)
	RegisterJoinHandler(t data.Type, handle JoinHandler)
	Start()
	GetClientByKey(key string) client.Client
}
