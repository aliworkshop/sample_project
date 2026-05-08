package delivery

import (
	"github.com/aliworkshop/errors"
	"github.com/aliworkshop/gateway/v2"
)

type hiHandler struct {
}

func NewHiHandler() gateway.Handler {
	handler := new(hiHandler)
	return handler
}

func (h *hiHandler) Handle(request gateway.HttpRequester) (any, errors.ErrorModel) {
	request.Paginator().SetTotal(1)
	return map[string]any{
		"message": "hello world",
	}, nil
}