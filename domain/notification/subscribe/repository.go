package subscribe

type SubscribeRepo interface {
	ListSubscribeTopicMappingBySubscriber(subscriber *Subscriber) ([]*SubscribeTopicMapping, error)
	SaveSubscribeTopicMapping(mapping *SubscribeTopicMapping) error
	RemoveSubscribeTopicMapping(mapping *SubscribeTopicMapping) error
}
