package delivery

import (
	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2"
)

type hiHandler struct {
}

func NewHiHandler() gateway.Handler {
	handler := new(hiHandler)
	return handler
}

func (h *hiHandler) Handle(request gateway.Requester) (any, errors.ErrorModel) {

	//fmt.Println("user", request.GetAuth().GetClaim().GetUserId())
	request.Paginator().SetTotal(1)
	return map[string]any{
		"message": "hello world",
	}, nil
}
