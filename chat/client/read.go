package client

import (
	"context"
	"encoding/json"

	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat/client/data"
	"github.com/gorilla/websocket"
)

func (c *client) read() {
	for {
		t, b, err := c.conn.Read(context.Background())
		if err != nil {
			// If Stop() already fired, the read error is expected (the conn
			// was closed under us); return quietly.
			select {
			case <-c.closed:
				return
			default:
			}

			c.log.
				With(logger.Field{
					"error": err,
				}).
				ErrorF("error on read message")
			if e, ok := err.(*websocket.CloseError); ok {
				if e.Code >= websocket.CloseNormalClosure && e.Code <= websocket.CloseTLSHandshake {
					c.Stop()
				}
			}
			return
		}
		switch t {
		case websocket.TextMessage:
			go c.handleTextMsg(b)
		}
	}
}

func (c *client) handleTextMsg(msg []byte) {
	req := new(data.Data)
	if err := json.Unmarshal(msg, req); err != nil {
		c.log.WarnF("error on unmarshal message: %v", err.Error())
		return
	}
	switch req.Action {
	case data.Message:
		c.handleMessage(req)
	case data.Join:
		c.handleJoin(req)
	}
}
