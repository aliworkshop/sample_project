package client

import (
	"fmt"
	"sync"
	"time"

	"github.com/aliworkshop/errors"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat/client/data"
	"github.com/aliworkshop/sample_project/chat/client/event"
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

	closed    chan struct{}
	stopOnce  sync.Once
	startOnce sync.Once

	writeChan chan *WriteRequest
	eventChan chan *Event

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
		closed:    make(chan struct{}),
		writeChan: make(chan *WriteRequest),
		eventChan: eventChan,
		key:       fmt.Sprintf("%d", userId),
		values:    make(map[string]interface{}),
		valuesMtx: new(sync.RWMutex),
	}
	conn.SetWriteDeadLine(5 * time.Minute)
	conn.SetReadDeadLine(5 * time.Minute)
	return c
}

func (c *client) Start() {
	c.startOnce.Do(func() {
		go c.read()
		go c.write()
	})
}

func (c *client) Stop() {
	c.stopOnce.Do(func() {
		c.connMtx.Lock()
		// Signal readers/writers to wind down before tearing the conn down.
		close(c.closed)
		close(c.writeChan)
		c.conn.Close()
		c.connMtx.Unlock()

		// Done outside the lock: it's a synchronous send into the chat use-case
		// loop and we mustn't hold the lock while the receiver runs handlers.
		c.eventChan <- &Event{
			Client: c,
			Event:  event.Closed,
		}
	})
}

func (c *client) IsAlive() bool {
	select {
	case <-c.closed:
		return false
	default:
		return true
	}
}

func (c *client) GetKey() string {
	return c.key
}
