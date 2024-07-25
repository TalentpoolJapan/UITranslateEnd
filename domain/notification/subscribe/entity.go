package subscribe

import (
	"uitranslate/domain/notification"
	"uitranslate/domain/notification/topic"
)

type SubscriberType string

const (
	JobSeeker  = SubscriberType("job_seeker")
	Enterprise = SubscriberType("enterprise")
)

type Subscriber struct {
	Uuid  string         `json:"uuid"`
	Type  SubscriberType `json:"type"`
	Name  string         `json:"name"`
	Email string         `json:"email"`
}

func (s Subscriber) AcceptChannels(topicInfo topic.TopicInfo) []notification.Channel {
	// todo
	return nil
}

type SubscribeTopicMapping struct {
	ID             int64          `json:"id"`
	SubscriberType SubscriberType `json:"subscriber_type"`
	SubscriberUuid string         `json:"subscriber_uuid"`
	TopicId        int64          `json:"topic_id"`
}

type SubscriberSubscribedTopic struct {
	Subscriber Subscriber
	TopicIds   []int64
}

func (s *SubscriberSubscribedTopic) alreadySubscribe(topicId int64) bool {
	for _, id := range s.TopicIds {
		if id == topicId {
			return true
		}
	}
	return false

}
