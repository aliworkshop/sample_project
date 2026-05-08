package client

import (
	"context"
	"time"

	"github.com/aliworkshop/errors"
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat/client/data"
	"github.com/gorilla/websocket"
)

type WriteRequest struct {
	Type int
	Data []byte
}

func (c *client) Write(w *WriteRequest) errors.ErrorModel {
	return c.writeMessage(w.Type, w.Data)
}

func (c *client) WriteBinary(data []byte) errors.ErrorModel {
	return c.writeMessage(websocket.BinaryMessage, data)
}

func (c *client) WriteJson(data *data.Data) errors.ErrorModel {
	c.connMtx.Lock()
	defer c.connMtx.Unlock()

	if err := c.conn.WriteJson(context.Background(), data.Body); err != nil {
		return errors.HandleError(err)
	}
	return nil
}

func (c *client) writeMessage(messageType int, data []byte) errors.ErrorModel {
	c.connMtx.Lock()
	defer c.connMtx.Unlock()

	if c.conn == nil {
		return errors.Internal().
			WithMessage("connection is closed")
	}

	if err := c.conn.Write(context.Background(), messageType, data); err != nil {
		return errors.HandleError(err)
	}
	return nil
}

func (c *client) write() {
	t := time.NewTicker(time.Second * 5)
	defer t.Stop()

	for c.IsAlive() {
		select {
		case w := <-c.writeChan:
			if w == nil {
				return
			}
			if err := c.writeMessage(w.Type, w.Data); err != nil {
				c.log.
					With(logger.Field{
						"error": err,
					}).
					WithId("c.conn.WriteMessage").
					ErrorF("error on conn.WriteMessage")
			}
		case <-t.C:
			err := c.conn.Write(context.Background(), websocket.PingMessage, nil)
			if err != nil {
				c.Stop()
				c.log.WithId("c.conn.PingMessage").With(
					logger.Field{
						"error": err.Error(),
					}).ErrorF("")
			}
		}
	}
}
