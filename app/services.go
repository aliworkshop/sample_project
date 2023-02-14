package app

import "github.com/aliworkshop/oauthlib"

func (a *App) initOauth() {
	a.oauth = oauthlib.NewAuthorizationHandler(a.initHr(), a.mainLogger)
}
