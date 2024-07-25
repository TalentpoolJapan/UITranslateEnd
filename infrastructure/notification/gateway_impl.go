package notification

import (
	"uitranslate/domain/notification/model"
	"uitranslate/domain/notification/topic"
	"uitranslate/infrastructure/notification/repo"
)

type GatewayImpl struct {
	notificationRepo repo.NotificationRepository
}

func newGatewayImpl(r repo.NotificationRepository) *GatewayImpl {
	return &GatewayImpl{
		notificationRepo: r,
	}
}

func (g *GatewayImpl) GetTopicInfoById(topicId int64) (*topic.TopicInfo, error) {
	topicInfo, err := g.notificationRepo.GetTopicInfoById(topicId)
	if err != nil {
		return nil, err
	}
	return topicInfo, nil
}

func (g *GatewayImpl) ListTopicInfo() ([]*topic.TopicInfo, error) {
	topicInfos, err := g.notificationRepo.ListTopicInfo()
	if err != nil {
		return nil, err
	}
	return topicInfos, nil
}

func (g *GatewayImpl) SaveTopicInfo(topicInfo *topic.TopicInfo) error {
	return g.notificationRepo.SaveTopicInfo(*topicInfo)
}

func (g *GatewayImpl) UpdateTopicInfo(topicInfo *topic.TopicInfo) error {
	return g.notificationRepo.UpdateTopicInfo(*topicInfo)
}

func (g *GatewayImpl) GetTopicTemplateById(templateId int64) (*topic.TopicTemplate, error) {
	return g.notificationRepo.GetTopicTemplateById(templateId)
}

func (g *GatewayImpl) ListTopicTemplateByTopicId(topicId int64) ([]*topic.TopicTemplate, error) {
	if topicId == 0 {
		return g.notificationRepo.ListTopicTemplate()
	} else {
		return g.notificationRepo.ListTopicTemplateByTopicId(topicId)
	}
}

func (g *GatewayImpl) SaveTopicTemplate(topicTemplate *topic.TopicTemplate) error {
	return g.notificationRepo.SaveTopicTemplate(*topicTemplate)
}

func (g *GatewayImpl) UpdateTopicTemplate(topicTemplate *topic.TopicTemplate) error {
	return g.notificationRepo.UpdateTopicTemplate(*topicTemplate)
}

func (g *GatewayImpl) GetAggregateTopicById(topicId int64) (*topic.AggregateTopic, error) {
	// todo
	return nil, nil
}

func (g *GatewayImpl) GetTriggerById(triggerId int64) (*model.Trigger, error) {
	return g.notificationRepo.GetTriggerById(triggerId)
}

func (g *GatewayImpl) ListTrigger() ([]*model.Trigger, error) {
	return g.notificationRepo.ListTrigger()
}

func (g *GatewayImpl) ListTriggerByTopicId(topicId int64) ([]*model.Trigger, error) {
	return g.notificationRepo.ListTrigger()
}

func (g *GatewayImpl) SaveTrigger(trigger *model.Trigger) error {
	return g.notificationRepo.SaveTrigger(*trigger)
}

func (g *GatewayImpl) UpdateTrigger(trigger *model.Trigger) error {
	return g.notificationRepo.UpdateTrigger(*trigger)
}
