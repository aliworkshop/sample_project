package client

import (
	"encoding/json"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/aliworkshop/sample_project/chat/client/data"
	"github.com/gorilla/websocket"
)

func (c *client) read() {
	for c.conn != nil {
		t, b, err := c.conn.Read()
		if err != nil {
			c.log.
				With(logger.Field{
					"error": err,
				}).
				ErrorF("error on read message")
			switch e := err.(type) {
			case *websocket.CloseError:
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
