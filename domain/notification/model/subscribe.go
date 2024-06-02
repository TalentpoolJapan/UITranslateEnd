package model

type Subscriber struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (s Subscriber) AcceptChannels(topicInfo TopicInfo) []Channel {
	// todo
	return nil
}

type SubscribeTopicMapping struct {
	SubscriberId int64 `json:"subscriber_id"`
	TopicId      int64 `json:"topic_id"`
}

type SubscribeTopic struct {
	Subscriber Subscriber
	TopicIds   []int64
}
