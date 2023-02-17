package delivery

import (
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/oauthlib/claim/domain"
	hd "github.com/aliworkshop/oauthlib/handler/domain"
	"github.com/golang-jwt/jwt/v4"
)

type loginHandler struct {
	handlerlib.HandlerModel
	oauth hd.Handler
}

func NewLoginHandler(handlerModel handlerlib.HandlerModel, oauth hd.Handler) handlerlib.HandlerModel {
	handler := new(loginHandler)
	handler.HandlerModel = handlerModel
	handler.oauth = oauth
	handler.SetHandlerFunc(handler.handle)
	return handler
}

func (h *loginHandler) handle(request handlerlib.RequestModel, args ...interface{}) (interface{}, errorslib.ErrorModel) {
	claims := domain.JWTClaim{
		Name:   "Ali torabi",
		Email:  "sralitorabi@gmail.com",
		Mobile: "09194768827",
		Scopes: []string{
			"api.project.login",
			"api.project.get",
			"api.project.paginate",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ID: "15",
		},
	}

	accessToken, refreshToken, err := h.oauth.GenerateTokens(claims)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}
