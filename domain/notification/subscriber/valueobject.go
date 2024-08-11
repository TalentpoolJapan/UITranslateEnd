package subscriber

type SubscribeTopicMapping struct {
	ID             int64  `json:"id"`
	SubscriberType Type   `json:"subscriber_type"`
	SubscriberUuid string `json:"subscriber_uuid"`
	TopicId        int64  `json:"topic_id"`
}
