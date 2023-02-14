package delivery

import (
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/oauthlib/claim/domain"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

type loginHandler struct {
	handlerlib.HandlerModel
}

func NewLoginHandler(handlerModel handlerlib.HandlerModel) handlerlib.HandlerModel {
	handler := new(loginHandler)
	handler.HandlerModel = handlerModel
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			ID:        "15",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, errorslib.Internal(err)
	}

	return map[string]any{
		"token": t,
	}, nil
}
