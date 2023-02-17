package delivery

import (
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib"
	cd "github.com/aliworkshop/oauthlib/claim/domain"
	hd "github.com/aliworkshop/oauthlib/handler/domain"
	"github.com/aliworkshop/sample_project/hello/domain"
	"github.com/golang-jwt/jwt/v4"
)

type refreshHandler struct {
	handlerlib.HandlerModel
	oauth hd.Handler
}

func NewRefreshHandler(handlerModel handlerlib.HandlerModel, oauth hd.Handler) handlerlib.HandlerModel {
	handler := new(refreshHandler)
	handler.HandlerModel = handlerModel
	handler.oauth = oauth
	handler.SetHandlerFunc(handler.handle)
	return handler
}

func (h *refreshHandler) handle(request handlerlib.RequestModel, args ...interface{}) (interface{}, errorslib.ErrorModel) {
	var req domain.RefreshRequest
	if err := request.HandleRequestBody(&req); err != nil {
		return nil, err
	}

	fetchedClaim, err := h.oauth.FetchTokenClaims(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	claims := cd.JWTClaim{
		Name:  "Ali torabi",
		Email: "sralitorabi@gmail.com",
		Scopes: []string{
			"api.project.login",
			"api.project.get",
			"api.project.paginate",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ID: fetchedClaim.ID,
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
