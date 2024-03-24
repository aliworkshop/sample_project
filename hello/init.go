package hello

import (
	"github.com/aliworkshop/authorizer/handler/domain"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/sample_project/hello/delivery"
)

type Module struct {
	Hi      gateway.Handler
	Post    gateway.Handler
	Login   gateway.Handler
	Refresh gateway.Handler
}

func New(oauth domain.Handler) *Module {
	m := new(Module)
	m.Hi = delivery.NewHiHandler()
	m.Post = delivery.NewPostHandler()
	m.Login = delivery.NewLoginHandler(oauth)
	m.Refresh = delivery.NewRefreshHandler(oauth)
	return m
}
