package gateway

import "uitranslate/domain/notification/model"

type Gateway interface {
	// topic
	GetTopicInfoById(topicId int64) (*model.TopicInfo, error)
	ListTopicInfo() ([]*model.TopicInfo, error)
	SaveTopicInfo(topicInfo *model.TopicInfo) error
	UpdateTopicInfo(topicInfo *model.TopicInfo) error

	// topic template
	GetTopicTemplateById(templateId int64) (*model.TopicTemplate, error)
	ListTopicTemplateByTopicId(topicId int64) ([]*model.TopicTemplate, error)
	SaveTopicTemplate(topicTemplate *model.TopicTemplate) error
	UpdateTopicTemplate(topicTemplate *model.TopicTemplate) error

	// aggregate topic
	GetAggregateTopicById(topicId int64) (*model.AggregateTopic, error)

	// trigger
	GetTriggerById(triggerId int64) (*model.Trigger, error)
	ListTriggerByTopicId(topicId int64) ([]*model.Trigger, error)
	ListTrigger() ([]*model.Trigger, error)
	SaveTrigger(trigger *model.Trigger) error
	UpdateTrigger(trigger *model.Trigger) error

	// subscribe
	SubscribeTopic(subscribeTopic *model.SubscribeTopic) error
	//ListSubscribeTopic(subscriber *model.Subscriber) (model.SubscribeTopic, error)
}
