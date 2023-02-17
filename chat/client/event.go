package client

import "github.com/aliworkshop/sample_project/chat/client/event"

type Event struct {
	Client Client
	Event  event.Event
}
