package delivery

import (
	"fmt"
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/sample_project/hello/domain"
)

type postHandler struct {
	handlerlib.HandlerModel
}

func NewPostHandler(handlerModel handlerlib.HandlerModel) handlerlib.HandlerModel {
	handler := new(postHandler)
	handler.HandlerModel = handlerModel
	handler.SetHandlerFunc(handler.handle)
	return handler
}

func (h *postHandler) handle(request handlerlib.RequestModel, args ...interface{}) (interface{}, errorslib.ErrorModel) {
	var req domain.PostRequest
	if err := request.HandleRequestBody(&req); err != nil {
		return nil, err
	}

	fmt.Println(request.GetParam("id"))
	return req, nil
}
