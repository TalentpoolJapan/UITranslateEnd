package http

import (
	"uitranslate/domain/notification/subscriber"
)

type httpUserGateway struct {
}

func NewHttpUserGateway() subscriber.IGateway {
	return &httpUserGateway{}
}

func (h httpUserGateway) GetJobseekerInfo(uuid string) (*subscriber.ExternalJobseekerInfo, error) {
	//TODO implement me
	panic("implement me")
}
