package app

import (
	"github.com/aliworkshop/gateway/v2"
)

func (a *App) RegisterRoutes() {
	mon := a.Monitor

	rg := a.engine.NewRouterGroup("/")
	gateway.RegisterRouters(rg, "hi", gateway.Read,
		mon.Wrap("hello", "hi", a.HiModule.Hi))
	gateway.RegisterRouters(rg, "post/:id", gateway.Create,
		a.oauth.MustHaveScope("api.project.get"),
		mon.Wrap("hello", "post", a.HiModule.Post))
	gateway.RegisterRouters(rg, "login", gateway.Create,
		mon.Wrap("hello", "login", a.HiModule.Login))
	gateway.RegisterRouters(rg, "refresh", gateway.Create,
		a.oauth.ShouldAuthenticate(),
		mon.Wrap("hello", "refresh", a.HiModule.Refresh))

	wsRg := a.engine.NewRouterGroup("/ws/")
	gateway.RegisterRouters(wsRg, "subscribe", gateway.Read,
		a.oauth.MustAuthenticate(),
		mon.Wrap("chat", "subscribe", a.ChatModule.Subscribe))
}
