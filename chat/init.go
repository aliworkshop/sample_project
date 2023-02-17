package chat

import (
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/aliworkshop/sample_project/chat/delivery"
	"github.com/aliworkshop/sample_project/chat/domain"
	"github.com/aliworkshop/sample_project/chat/usecase"
)

type Module struct {
	Uc        domain.ChatUc
	Subscribe handlerlib.HandlerModel
}

func New(model func() handlerlib.HandlerModel, logger logger.Logger) *Module {
	m := new(Module)
	m.Uc = usecase.NewUseCase(logger)
	m.Subscribe = delivery.NewSubscribeHandler(model(), m.Uc)
	return m
}
