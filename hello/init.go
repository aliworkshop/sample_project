package hello

import (
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/oauthlib/handler/domain"
	"github.com/aliworkshop/sample_project/hello/delivery"
)

type Module struct {
	Hi      handlerlib.HandlerModel
	Post    handlerlib.HandlerModel
	Login   handlerlib.HandlerModel
	Refresh handlerlib.HandlerModel
}

func New(model func() handlerlib.HandlerModel, oauth domain.Handler) *Module {
	m := new(Module)
	m.Hi = delivery.NewHiHandler(model())
	m.Post = delivery.NewPostHandler(model())
	m.Login = delivery.NewLoginHandler(model(), oauth)
	m.Refresh = delivery.NewRefreshHandler(model(), oauth)
	return m
}
