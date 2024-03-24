package delivery

import (
	"fmt"
	cd "github.com/aliworkshop/authorizer/claim/domain"
	hd "github.com/aliworkshop/authorizer/handler/domain"
	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type refreshHandler struct {
	oauth hd.Handler
}

func NewRefreshHandler(oauth hd.Handler) gateway.Handler {
	handler := new(refreshHandler)
	handler.oauth = oauth
	return handler
}

func (h *refreshHandler) Handle(request gateway.Requester) (any, errors.ErrorModel) {
	if request.GetAuth() == nil {
		return nil, errors.DefaultUnAuthenticatedError
	}

	jwtClaim := cd.JWTClaim{
		Uuid:             uuid.NewString(),
		RegisteredClaims: jwt.RegisteredClaims{ID: fmt.Sprintf("%d", request.GetCurrentAccountId())},
	}

	claim := cd.Claim{
		UserId: request.GetCurrentAccountId(),
		Name:   request.GetAuth().GetClaim().GetName(),
		Email:  request.GetAuth().GetClaim().GetEmail(),
		Mobile: request.GetAuth().GetClaim().GetMobile(),
		Scopes: request.GetAuth().GetClaim().GetScopes(),
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
