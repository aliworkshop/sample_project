package client

import (
	"fmt"
	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat/client/data"
	"github.com/aliworkshop/sample_project/chat/client/event"
	"sync"
	"time"
)

type Client interface {
	Start()
	Stop()
	IsAlive() bool

	GetKey() string

	SetTemp(key string, value interface{})
	GetTemp(key string) interface{}

	Write(w *WriteRequest) errors.ErrorModel
	WriteBinary(data []byte) errors.ErrorModel
	WriteJson(data *data.Data) errors.ErrorModel
}

type client struct {
	log logger.Logger

	conn    gateway.WebSocketHandler
	connMtx *sync.Mutex

	writeChan chan *WriteRequest
	eventChan chan *Event

	started bool
	// key unique key of connection
	key string

	values    map[string]interface{}
	valuesMtx *sync.RWMutex
}

func New(log logger.Logger, conn gateway.WebSocketHandler, userId uint64, eventChan chan *Event) Client {
	c := &client{
		log:       log,
		conn:      conn,
		connMtx:   new(sync.Mutex),
		writeChan: make(chan *WriteRequest),
		eventChan: eventChan,
		key:       fmt.Sprintf("%d", userId),
		values:    make(map[string]interface{}),
		valuesMtx: new(sync.RWMutex),
	}
	conn.SetWriteDeadLine(5 * time.Minute)
	conn.SetReadDeadLine(5 * time.Minute)
	conn.SetCloseHandler(c.closeHandler)
	return c
}

func (c *client) Start() {
	if c.started {
		return
	}
	c.started = true

	go c.read()
	go c.write()
}

func (c *client) Stop() {
	if !c.started {
		return
	}

	c.connMtx.Lock()
	defer c.connMtx.Unlock()

	close(c.writeChan)

	c.conn.Close()
	c.conn = nil

	c.eventChan <- &Event{
		Client: c,
		Event:  event.Closed,
	}

	c.started = false
}

func (c *client) IsAlive() bool {
	return c.conn != nil
}

func (c *client) GetKey() string {
	return c.key
}
