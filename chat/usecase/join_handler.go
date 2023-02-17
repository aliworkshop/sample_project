package usecase

import (
	"github.com/aliworkshop/sample_project/chat/client"
	"github.com/aliworkshop/sample_project/chat/client/data"
)

func (uc *useCase) HandleJoinGroup(c client.Client, request *data.Data) {
	if group, ok := uc.groups[request.Id]; ok {
		if !keyExistsInArray(c.GetKey(), group.Members) {
			group.Members = append(group.Members, c.GetKey())
			uc.groups[request.Id] = group
		}
	}
}

func (uc *useCase) HandleJoinChannel(c client.Client, request *data.Data) {
	if channel, ok := uc.channels[request.Id]; ok {
		if !keyExistsInArray(c.GetKey(), channel.Members) {
			channel.Members = append(channel.Members, c.GetKey())
			uc.channels[request.Id] = channel
		}
	}
}
