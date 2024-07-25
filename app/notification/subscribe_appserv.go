package notification

type SubscribeTopicCmd struct {
	SubscriberUuid string
	SubscriberType string
	TopicId        int64
}

type UnsubscribeTopicCmd struct {
	SubscriberUuid string
	SubscriberType string
	TopicId        int64
}

type SubscribeAppServ interface {
	SubscribeTopic(cmd SubscribeTopicCmd) error
	UnsubscribeTopic(cmd UnsubscribeTopicCmd) error
}
