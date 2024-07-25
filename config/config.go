package config

import (
	"uitranslate/app/notification"
	nf_gateway "uitranslate/domain/notification/gateway"
)

var (
	subscribeAppServ *notification.SubscribeAppServ
	topicAppServ     *notification.TopicAppServ

	notificationGateway nf_gateway.Gateway
	topicGateway        nf_gateway.TopicGateway
	userGateway         nf_gateway.UserGateway
)
