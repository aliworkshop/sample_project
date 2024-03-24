package usecase

import (
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat/client"
	"github.com/aliworkshop/sample_project/chat/client/data"
	"github.com/aliworkshop/sample_project/chat/client/event"
	"github.com/aliworkshop/sample_project/chat/domain"
)

type useCase struct {
	started bool

	eventChan       chan *client.Event
	requestHandlers map[data.Type][]domain.RequestHandle
	joinHandlers    map[data.Type][]domain.JoinHandler

	clients  map[string]client.Client
	groups   map[string]domain.Group
	channels map[string]domain.Channel
	logger   logger.Logger
}

func NewUseCase(logger logger.Logger) domain.ChatUc {
	uc := &useCase{
		logger:          logger,
		eventChan:       make(chan *client.Event),
		requestHandlers: make(map[data.Type][]domain.RequestHandle),
		joinHandlers:    make(map[data.Type][]domain.JoinHandler),
		clients:         make(map[string]client.Client),
		groups:          make(map[string]domain.Group),
		channels:        make(map[string]domain.Channel),
	}
	uc.groups["1"] = domain.Group{
		Name: "first group",
		Id:   "1",
	}
	uc.channels["1"] = domain.Channel{
		Name:  "first channel",
		Id:    "1",
		Admin: "15",
	}
	uc.RegisterRequestHandler(data.User, uc.HandlerPrivateMsg)
	uc.RegisterRequestHandler(data.Group, uc.HandlerGroupMsg)
	uc.RegisterRequestHandler(data.Channel, uc.HandlerChannelMsg)

	uc.RegisterJoinHandler(data.Group, uc.HandleJoinGroup)
	uc.RegisterJoinHandler(data.Channel, uc.HandleJoinChannel)
	return uc
}

func (uc *useCase) RegisterRequestHandler(t data.Type, handle domain.RequestHandle) {
	if uc.started {
		panic("can not register handler while websocket is started")
	}
	handlers := uc.requestHandlers[t]
	if handlers == nil {
		handlers = make([]domain.RequestHandle, 0)
	}
	handlers = append(handlers, handle)
	uc.requestHandlers[t] = handlers
}

func (uc *useCase) RegisterJoinHandler(t data.Type, handle domain.JoinHandler) {
	if uc.started {
		panic("can not register handler while websocket is started")
	}
	handlers := uc.joinHandlers[t]
	if handlers == nil {
		handlers = make([]domain.JoinHandler, 0)
	}
	handlers = append(handlers, handle)
	uc.joinHandlers[t] = handlers
}

func (uc *useCase) Start() {
	if uc.started {
		return
	}
	uc.started = true
	go uc.start()
}

func (uc *useCase) GetClientByKey(key string) client.Client {
	return uc.clients[key]
}

func (uc *useCase) start() {
	for uc.started {
		select {
		case e := <-uc.eventChan:
			if e == nil {
				return
			}
			switch e.Event.Type {
			case event.TypeMessage:
				handlers := uc.requestHandlers[e.Event.Request.Type]
				for _, h := range handlers {
					go h(e.Client, e.Event.Request)
				}
				break
			case event.TypeClosed:
				delete(uc.clients, e.Client.GetKey())
				break
			case event.TypeJoin:
				handlers := uc.joinHandlers[e.Event.Request.Type]
				for _, h := range handlers {
					go h(e.Client, e.Event.Request)
				}
				break
			}
			break
		}
	}
}
