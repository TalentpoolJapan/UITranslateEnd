package subscriber

type IRepository interface {
	ListSubscribedTopicIdsBySubscriber(userType Type, userUuid string) ([]int64, error)
	SaveSubscriberTopic(subscriber *Subscriber) error
}
