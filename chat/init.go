package chat

import (
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat/delivery"
	"github.com/aliworkshop/sample_project/chat/domain"
	"github.com/aliworkshop/sample_project/chat/usecase"
)

type Module struct {
	Uc        domain.ChatUc
	Subscribe gateway.Handler
}

func New(logger logger.Logger) *Module {
	m := new(Module)
	m.Uc = usecase.NewUseCase(logger)
	m.Subscribe = delivery.NewSubscribeHandler(m.Uc)
	return m
}
