package config

import (
	"uitranslate/app/notification"
	"uitranslate/domain/notification/subscriber"
	"uitranslate/domain/notification/topic"
	nf_gateway "uitranslate/domain/notification/trigger"
)

var (
	subscribeAppServ *notification.SubscribeAppServ
	topicAppServ     *notification.TopicAppServ

	notificationGateway nf_gateway.Repository
	topicGateway        topic.Repository
	userGateway         subscriber.IGateway
)
