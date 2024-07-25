package subscribe

import (
	"log"
)

type SubscribeDomainService interface {
	SubscribeTopic(subscriber *Subscriber, topicId int64) error
	UnsubscribeTopic(subscriber *Subscriber, topicId int64) error
	FindSubscribeTopic(subscriber *Subscriber) (*SubscriberSubscribedTopic, error)
}

type subscribeDomainService struct {
	subscribeRepo SubscribeRepo
}

func NewSubscribeDomainService(subscribeRepo SubscribeRepo) SubscribeDomainService {
	return &subscribeDomainService{
		subscribeRepo: subscribeRepo,
	}
}

func (g *subscribeDomainService) SubscribeTopic(subscriber *Subscriber, topicId int64) error {
	subscribedTopic, err := g.FindSubscribeTopic(subscriber)
	if err != nil {
		return err
	}
	if subscribedTopic.alreadySubscribe(topicId) {
		log.Printf("Subscriber %s already subscribe topic %d", subscriber.Uuid, topicId)
		return nil
	}
	mapping := &SubscribeTopicMapping{
		SubscriberType: subscriber.Type,
		SubscriberUuid: subscriber.Uuid,
		TopicId:        topicId,
	}
	err = g.subscribeRepo.SaveSubscribeTopicMapping(mapping)
	if err != nil {
		return err
	}
	return nil
}

func (g *subscribeDomainService) UnsubscribeTopic(subscriber *Subscriber, topicId int64) error {
	subscribedTopic, err := g.FindSubscribeTopic(subscriber)
	if err != nil {
		return err
	}
	if !subscribedTopic.alreadySubscribe(topicId) {
		log.Printf("Subscriber %s already unsubscribe topic %d", subscriber.Uuid, topicId)
		return nil
	}
	mapping := &SubscribeTopicMapping{
		SubscriberType: subscriber.Type,
		SubscriberUuid: subscriber.Uuid,
		TopicId:        topicId,
	}
	err = g.subscribeRepo.RemoveSubscribeTopicMapping(mapping)
	if err != nil {
		return err
	}
	return nil
}

func (g *subscribeDomainService) FindSubscribeTopic(subscriber *Subscriber) (*SubscriberSubscribedTopic, error) {
	subscribeTopicMappings, err := g.subscribeRepo.ListSubscribeTopicMappingBySubscriber(subscriber)
	if err != nil {
		return nil, err
	}
	var topicIds []int64
	for _, mapping := range subscribeTopicMappings {
		topicIds = append(topicIds, mapping.TopicId)
	}
	return &SubscriberSubscribedTopic{
		Subscriber: *subscriber,
		TopicIds:   topicIds,
	}, nil
}
