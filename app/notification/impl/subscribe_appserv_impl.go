package impl

import (
	"errors"
	"log"
	"uitranslate/app/notification"
	"uitranslate/domain/notification/gateway"
	"uitranslate/domain/notification/subscriber"
	"uitranslate/domain/notification/topic"
)

type SubscribeAppServImpl struct {
	userGateway         gateway.UserGateway
	topicGateway        gateway.TopicGateway
	subscribeDomainServ subscriber.SubscribeDomainService
}

func NewSubscribeAppServImpl(userGateway gateway.UserGateway,
	topicGateway gateway.TopicGateway,
	subscribeDomainServ subscriber.SubscribeDomainService) notification.SubscribeAppServ {
	return &SubscribeAppServImpl{
		userGateway:         userGateway,
		topicGateway:        topicGateway,
		subscribeDomainServ: subscribeDomainServ,
	}
}

func (s *SubscribeAppServImpl) SubscribeTopic(cmd notification.SubscribeTopicCmd) error {
	userInfo, getUserErr := s.checkSubscriber(cmd.SubscriberUuid, cmd.SubscriberType)
	if getUserErr != nil {
		return getUserErr
	}

	_, getTopicInfoErr := s.checkTopicInfo(cmd.TopicId)
	if getTopicInfoErr != nil {
		return getTopicInfoErr
	}

	// subscriber
	subscriber := &subscriber.Subscriber{
		Uuid:  userInfo.Uuid,
		Name:  userInfo.Name,
		Email: userInfo.Email,
	}
	subscribeErr := s.subscribeDomainServ.SubscribeTopic(subscriber, cmd.TopicId)
	if subscribeErr != nil {
		return subscribeErr
	}
	return nil
}

func (s *SubscribeAppServImpl) UnsubscribeTopic(cmd notification.UnsubscribeTopicCmd) error {
	userInfo, getUserErr := s.checkSubscriber(cmd.SubscriberUuid, cmd.SubscriberType)
	if getUserErr != nil {
		return getUserErr
	}

	_, getTopicInfoErr := s.checkTopicInfo(cmd.TopicId)
	if getTopicInfoErr != nil {
		return getTopicInfoErr
	}

	// unsubscribe
	subscriber := &subscriber.Subscriber{
		Uuid:  userInfo.Uuid,
		Name:  userInfo.Name,
		Email: userInfo.Email,
	}
	subscribeErr := s.subscribeDomainServ.UnsubscribeTopic(subscriber, cmd.TopicId)
	if subscribeErr != nil {
		return subscribeErr
	}
	return nil
}

func (s *SubscribeAppServImpl) checkSubscriber(uuid string, userType string) (*gateway.ExternalUserInfo, error) {
	userInfo, getUserErr := s.userGateway.GetUserInfo(uuid, userType)
	if getUserErr != nil || userInfo == nil {
		log.Printf("%s-%s find user info error: %v", uuid, userType, getUserErr)
		return nil, errors.New("user not found")
	}
	return userInfo, nil
}

func (s *SubscribeAppServImpl) checkTopicInfo(topicId int64) (*topic.TopicInfo, error) {
	topicInfo, getTopicErr := s.topicGateway.GetTopicInfoById(topicId)
	if getTopicErr != nil || topicInfo == nil {
		log.Printf("%d find topic info", topicId)
		return nil, errors.New("topic not found")
	}
	return topicInfo, nil
}
