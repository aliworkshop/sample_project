package usecase

import (
	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/sample_project/chat/client"
)

func (uc *useCase) Subscribe(userId uint64, ws gateway.WebSocketHandler) (client.Client, errors.ErrorModel) {

	c := client.New(uc.logger, ws, userId, uc.eventChan)
	uc.clients[c.GetKey()] = c
	c.Start()

	return c, nil
}

func keyExistsInArray(key string, array []string) bool {
	for _, elm := range array {
		if elm == key {
			return true
		}
	}
	return false
}
