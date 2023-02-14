package delivery

import (
	"fmt"
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib"
)

type getHandler struct {
	handlerlib.HandlerModel
}

func NewHiHandler(handlerModel handlerlib.HandlerModel) handlerlib.HandlerModel {
	handler := new(getHandler)
	handler.HandlerModel = handlerModel
	handler.SetHandlerFunc(handler.handle)
	return handler
}

func (h *getHandler) handle(request handlerlib.RequestModel, args ...interface{}) (interface{}, errorslib.ErrorModel) {

	fmt.Println("user", request.GetAuth().GetClaim().GetUserId())
	return map[string]any{
		"message": "hello world",
	}, nil
}
