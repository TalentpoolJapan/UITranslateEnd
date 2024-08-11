package notification

import (
	"uitranslate/domain/notification/topic"
	"uitranslate/domain/notification/trigger"
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

func (g *GatewayImpl) GetTopicInfoById(topicId int64) (*topic.BasicInfo, error) {
	topicInfo, err := g.notificationRepo.GetTopicInfoById(topicId)
	if err != nil {
		return nil, err
	}
	return topicInfo, nil
}

func (g *GatewayImpl) ListTopicInfo() ([]*topic.BasicInfo, error) {
	topicInfos, err := g.notificationRepo.ListTopicInfo()
	if err != nil {
		return nil, err
	}
	return topicInfos, nil
}

func (g *GatewayImpl) SaveTopicInfo(topicInfo *topic.BasicInfo) error {
	return g.notificationRepo.SaveTopicInfo(*topicInfo)
}

func (g *GatewayImpl) UpdateTopicInfo(topicInfo *topic.BasicInfo) error {
	return g.notificationRepo.UpdateTopicInfo(*topicInfo)
}

func (g *GatewayImpl) GetTopicTemplateById(templateId int64) (*topic.Template, error) {
	return g.notificationRepo.GetTopicTemplateById(templateId)
}

func (g *GatewayImpl) ListTopicTemplateByTopicId(topicId int64) ([]*topic.Template, error) {
	if topicId == 0 {
		return g.notificationRepo.ListTopicTemplate()
	} else {
		return g.notificationRepo.ListTopicTemplateByTopicId(topicId)
	}
}

func (g *GatewayImpl) SaveTopicTemplate(topicTemplate *topic.Template) error {
	return g.notificationRepo.SaveTopicTemplate(*topicTemplate)
}

func (g *GatewayImpl) UpdateTopicTemplate(topicTemplate *topic.Template) error {
	return g.notificationRepo.UpdateTopicTemplate(*topicTemplate)
}

func (g *GatewayImpl) GetAggregateTopicById(topicId int64) (*topic.AggregateTopic, error) {
	// todo
	return nil, nil
}

func (g *GatewayImpl) GetTriggerById(triggerId int64) (*trigger.Trigger, error) {
	return g.notificationRepo.GetTriggerById(triggerId)
}

func (g *GatewayImpl) ListTrigger() ([]*trigger.Trigger, error) {
	return g.notificationRepo.ListTrigger()
}

func (g *GatewayImpl) ListTriggerByTopicId(topicId int64) ([]*trigger.Trigger, error) {
	return g.notificationRepo.ListTrigger()
}

func (g *GatewayImpl) SaveTrigger(trigger *trigger.Trigger) error {
	return g.notificationRepo.SaveTrigger(*trigger)
}

func (g *GatewayImpl) UpdateTrigger(trigger *trigger.Trigger) error {
	return g.notificationRepo.UpdateTrigger(*trigger)
}
