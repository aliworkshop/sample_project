package delivery

import (
	"fmt"
	errors "github.com/aliworkshop/error"
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

func (h *postHandler) Handle(request gateway.Requester) (any, errors.ErrorModel) {
	var req domain.PostRequest
	if err := request.BindRequest(&req); err != nil {
		return nil, err
	}

	fmt.Println(request.GetParam("id"))
	return req, nil
}
