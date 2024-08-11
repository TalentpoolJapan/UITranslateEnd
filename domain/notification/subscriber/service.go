package subscriber

import (
	"log"
)

type SubscribeDomainService interface {
	SubscribeTopic(subscriber *Subscriber, topicId int64) error
	UnsubscribeTopic(subscriber *Subscriber, topicId int64) error
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
	if subscriber.alreadySubscribe(topicId) {
		log.Printf("Subscriber %s already subscriber topic %d", subscriber.Uuid, topicId)
		return nil
	}
	mapping := &SubscribeTopicMapping{
		SubscriberType: subscriber.Type,
		SubscriberUuid: subscriber.Uuid,
		TopicId:        topicId,
	}
	err := g.subscribeRepo.SaveSubscribeTopicMapping(mapping)
	if err != nil {
		return err
	}
	return nil
}

func (g *subscribeDomainService) UnsubscribeTopic(subscriber *Subscriber, topicId int64) error {
	if !subscriber.alreadySubscribe(topicId) {
		log.Printf("Subscriber %s already unsubscribe topic %d", subscriber.Uuid, topicId)
		return nil
	}
	mapping := &SubscribeTopicMapping{
		SubscriberType: subscriber.Type,
		SubscriberUuid: subscriber.Uuid,
		TopicId:        topicId,
	}
	err := g.subscribeRepo.RemoveSubscribeTopicMapping(mapping)
	if err != nil {
		return err
	}
	return nil
}
