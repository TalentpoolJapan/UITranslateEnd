package notification

type Subscriber struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SubscribeTopicMapping struct {
	SubscriberId int64 `json:"subscriber_id"`
	TopicId      int64 `json:"topic_id"`
}
