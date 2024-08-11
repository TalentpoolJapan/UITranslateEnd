package subscriber

import (
	"log"
)

type IDomainService interface {
	GetSubscriberByUniqueKey(subscriberType Type, subscriberOpenId string) (subscriber *Subscriber, err error)
	SubscribeTopic(subscriber *Subscriber, topicId int64) error
	UnsubscribeTopic(subscriber *Subscriber, topicId int64) error
	ListSubscriberByTopic(topicId int64) ([]*Subscriber, error)
}

type subscribeDomainService struct {
	subscribeRepo IRepository
}

func NewSubscribeDomainService(subscribeRepo IRepository) IDomainService {
	return &subscribeDomainService{
		subscribeRepo: subscribeRepo,
	}
}

func (g *subscribeDomainService) SubscribeTopic(subscriber *Subscriber, topicId int64) error {
	if !subscriber.subscribe(topicId) {
		log.Printf("Subscriber %s already subscriber topic %d", subscriber.Uuid, topicId)
		return nil
	}

	err := g.subscribeRepo.SaveSubscriberTopic(subscriber)
	if err != nil {
		return err
	}
	return nil
}

func (g *subscribeDomainService) UnsubscribeTopic(subscriber *Subscriber, topicId int64) error {
	if !subscriber.unsubscribe(topicId) {
		log.Printf("Subscriber %s already unsubscribe topic %d", subscriber.Uuid, topicId)
		return nil
	}
	err := g.subscribeRepo.SaveSubscriberTopic(subscriber)
	if err != nil {
		return err
	}
	return nil
}

func (g *subscribeDomainService) GetSubscriberByUniqueKey(subscriberType Type, subscriberOpenId string) (subscriber *Subscriber, err error) {

	//TODO implement me
	panic("implement me")
}

func (g *subscribeDomainService) ListSubscriberByTopic(topicId int64) ([]*Subscriber, error) {
	//TODO implement me
	panic("implement me")
}
