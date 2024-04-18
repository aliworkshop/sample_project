package app

import (
	"github.com/aliworkshop/gateway/v2"
)

func (a *App) RegisterRoutes() {
	rg := a.engine.NewRouterGroup("/")
	gateway.RegisterRouters(rg, "hi", gateway.Read, a.oauth.MustAuthenticate(), a.HiModule.Hi)
	gateway.RegisterRouters(rg, "post/:id", gateway.Create, a.oauth.MustHaveScope("api.project.get"), a.HiModule.Post)
	gateway.RegisterRouters(rg, "login", gateway.Create, a.HiModule.Login)
	gateway.RegisterRouters(rg, "refresh", gateway.Create, a.oauth.ShouldAuthenticate(), a.HiModule.Refresh)

	wsRg := a.engine.NewRouterGroup("/ws/")
	gateway.RegisterRouters(wsRg, "subscribe", gateway.Read, a.oauth.MustAuthenticate(), a.ChatModule.Subscribe)

}
