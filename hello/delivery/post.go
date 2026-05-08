package delivery

import (
	"fmt"

	"github.com/aliworkshop/errors"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/sample_project/hello/domain"
)

type postHandler struct {
	gateway.Handler
}

func NewPostHandler() gateway.Handler {
	handler := new(postHandler)
	return handler
}

func (h *postHandler) Handle(request gateway.HttpRequester) (any, errors.ErrorModel) {
	var req domain.PostRequest
	if err := request.BindRequest(&req); err != nil {
		return nil, err
	}

	fmt.Println(request.GetParam("id"))
	return req, nil
}