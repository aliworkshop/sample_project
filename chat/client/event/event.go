package event

import "github.com/aliworkshop/sample_project/chat/client/data"

type Event struct {
	Type    Type
	Request *data.Data
}

var Closed = Event{
	Type: TypeClosed,
}
