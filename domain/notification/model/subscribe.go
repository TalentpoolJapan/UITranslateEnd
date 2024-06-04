package model

type SubscriberType string

const (
	JobSeeker  = SubscriberType("job_seeker")
	Enterprise = SubscriberType("enterprise")
)

type Subscriber struct {
	Uuid string         `json:"uuid"`
	Type SubscriberType `json:"type"`
	Name string         `json:"name"`
}

func (s Subscriber) AcceptChannels(topicInfo TopicInfo) []Channel {
	// todo
	return nil
}

type SubscribeTopicMapping struct {
	SubscriberType SubscriberType `json:"subscriber_type"`
	SubscriberUuid string         `json:"subscriber_uuid"`
	TopicId        int64          `json:"topic_id"`
}

type SubscribeTopic struct {
	Subscriber Subscriber
	TopicIds   []int64
}
