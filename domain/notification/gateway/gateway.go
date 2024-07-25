package gateway

import (
	"uitranslate/domain/notification/model"
	"uitranslate/domain/notification/topic"
)

type Gateway interface {

	// aggregate topic
	GetAggregateTopicById(topicId int64) (*topic.AggregateTopic, error)

	// trigger
	GetTriggerById(triggerId int64) (*model.Trigger, error)
	ListTriggerByTopicId(topicId int64) ([]*model.Trigger, error)
	ListTrigger() ([]*model.Trigger, error)
	SaveTrigger(trigger *model.Trigger) error
	UpdateTrigger(trigger *model.Trigger) error
}
