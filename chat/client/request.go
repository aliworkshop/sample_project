package client

import (
	"github.com/aliworkshop/sample_project/chat/client/data"
	"github.com/aliworkshop/sample_project/chat/client/event"
)

func (c *client) handleMessage(req *data.Data) {
	c.eventChan <- &Event{
		Client: c,
		Event: event.Event{
			Type:    event.TypeMessage,
			Request: req,
		},
	}
}

func (c *client) handleJoin(req *data.Data) {
	c.eventChan <- &Event{
		Client: c,
		Event: event.Event{
			Type:    event.TypeJoin,
			Request: req,
		},
	}
}
