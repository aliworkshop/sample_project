package delivery

import (
	"github.com/aliworkshop/authorizer/claim/domain"
	hd "github.com/aliworkshop/authorizer/handler/domain"
	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type loginHandler struct {
	oauth hd.Handler
}

func NewLoginHandler(oauth hd.Handler) gateway.Handler {
	return &loginHandler{oauth: oauth}
}

func (h *loginHandler) Handle(request gateway.Requester) (any, errors.ErrorModel) {
	jwtClaim := domain.JWTClaim{
		Uuid:             uuid.NewString(),
		RegisteredClaims: jwt.RegisteredClaims{ID: "15"},
	}

	claim := domain.Claim{
		UserId: 15,
		Name:   "Ali Torabi",
		Email:  "sralitorabi@gmail.com",
		Mobile: "09194768827",
		Scopes: []string{
			"api.content.advertise.create",
			"api.content.advertise.update",
			"api.content.advertise.delete",
		},
	}

	accessToken, refreshToken, err := h.oauth.GenerateTokens(jwtClaim, claim)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}
