package app

import (
	"github.com/aliworkshop/echoserver"
	"github.com/aliworkshop/handlerlib"
)

func (a *App) RegisterRoutes() {
	rg := a.engine.NewRouterGroup("/")
	handlerlib.RegisterRouters(rg, "salam", handlerlib.ActionRead, a.oauth.MustAuthenticate(), a.HiModule.Hi)
	handlerlib.RegisterRouters(rg, "post/:id", handlerlib.ActionCreate, a.oauth.MustHaveScope("api.project.get"), a.HiModule.Post)
	handlerlib.RegisterRouters(rg, "login", handlerlib.ActionCreate, a.HiModule.Login)
	handlerlib.RegisterRouters(rg, "refresh", handlerlib.ActionCreate, a.HiModule.Refresh)

	wsRg := a.engine.NewRouterGroup("/ws/")
	handlerlib.RegisterRouters(wsRg, "subscribe", handlerlib.ActionRead, a.oauth.MustAuthenticate(), a.ChatModule.Subscribe)

}

func (a *App) initNoRespHr() (handlerModel handlerlib.HandlerModel) {
	responder := echoserver.NewEmptyResponder(nil)
	baseHandler := echoserver.NewHandler(a.mainLogger)
	return handlerlib.NewModel(baseHandler, responder)
}
