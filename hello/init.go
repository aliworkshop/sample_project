package hello

import (
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/sample_project/hello/delivery"
)

type Module struct {
	Hi    handlerlib.HandlerModel
	Post  handlerlib.HandlerModel
	Login handlerlib.HandlerModel
}

func New(model func() handlerlib.HandlerModel) *Module {
	m := new(Module)
	m.Hi = delivery.NewHiHandler(model())
	m.Post = delivery.NewPostHandler(model())
	m.Login = delivery.NewLoginHandler(model())
	return m
}
