package app

import "github.com/aliworkshop/authorizer"

func (a *App) initOauth() {
	a.oauth = authorizer.NewAuthorizationHandler(a.mainLogger, a.registry.ValueOf("oauth"))
}
