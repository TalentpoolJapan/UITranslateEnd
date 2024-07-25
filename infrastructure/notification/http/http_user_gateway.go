package http

import (
	"uitranslate/domain/notification/gateway"
)

type httpUserGateway struct {
}

func NewHttpUserGateway() gateway.UserGateway {
	return &httpUserGateway{}
}

func (h httpUserGateway) GetUserInfo(uuid string, userType string) (*gateway.ExternalUserInfo, error) {
	//TODO implement me
	panic("implement me")
}
