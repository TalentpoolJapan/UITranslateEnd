package notification

import (
	"uitranslate/domain/notification/model"
	"uitranslate/infrastructure/notification/repo"
)

var GatewaySingleton = NewGatewayImpl()

type GatewayImpl struct {
	notificationRepo repo.NotificationRepository
}

func NewGatewayImpl() *GatewayImpl {
	return &GatewayImpl{
		notificationRepo: *repo.Repo,
	}
}

func (g *GatewayImpl) GetTopicInfoById(topicId int64) (*model.TopicInfo, error) {
	topicInfo, err := g.notificationRepo.GetTopicInfoById(topicId)
	if err != nil {
		return nil, err
	}
	return topicInfo, nil
}

func (g *GatewayImpl) ListTopicInfo() ([]*model.TopicInfo, error) {
	topicInfos, err := g.notificationRepo.ListTopicInfo()
	if err != nil {
		return nil, err
	}
	return topicInfos, nil
}

func (g *GatewayImpl) SaveTopicInfo(topicInfo *model.TopicInfo) error {
	return g.notificationRepo.SaveTopicInfo(*topicInfo)
}

func (g *GatewayImpl) UpdateTopicInfo(topicInfo *model.TopicInfo) error {
	return g.notificationRepo.UpdateTopicInfo(*topicInfo)
}

func (g *GatewayImpl) GetTopicTemplateById(templateId int64) (*model.TopicTemplate, error) {
	return g.notificationRepo.GetTopicTemplateById(templateId)
}

func (g *GatewayImpl) ListTopicTemplateByTopicId(topicId int64) ([]*model.TopicTemplate, error) {
	return g.notificationRepo.ListTopicTemplateByTopicId(topicId)
}

func (g *GatewayImpl) SaveTopicTemplate(topicTemplate *model.TopicTemplate) error {
	return g.notificationRepo.SaveTopicTemplate(*topicTemplate)
}

func (g *GatewayImpl) UpdateTopicTemplate(topicTemplate *model.TopicTemplate) error {
	return g.notificationRepo.UpdateTopicTemplate(*topicTemplate)
}

func (g *GatewayImpl) GetAggregateTopicById(topicId int64) (*model.AggregateTopic, error) {
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

func (g *GatewayImpl) SubscribeTopic(subscribeTopic *model.SubscribeTopic) error {
	// todo
	return nil
}
