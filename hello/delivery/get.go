package delivery

import (
	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2"
)

type getHandler struct {
}

func NewHiHandler() gateway.Handler {
	handler := new(getHandler)
	return handler
}

func (h *getHandler) Handle(request gateway.Requester) (any, errors.ErrorModel) {

	//fmt.Println("user", request.GetAuth().GetClaim().GetUserId())
	request.Paginator().SetTotal(25)
	return map[string]any{
		"message": "hello world",
	}, nil
}
