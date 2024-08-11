package trigger

import (
	"uitranslate/domain/notification/topic"
)

type Repository interface {

	// aggregate topic
	GetAggregateTopicById(topicId int64) (*topic.AggregateTopic, error)

	// trigger
	GetTriggerById(triggerId int64) (*Trigger, error)
	ListTriggerByTopicId(topicId int64) ([]*Trigger, error)
	ListTrigger() ([]*Trigger, error)
	SaveTrigger(trigger *Trigger) error
	UpdateTrigger(trigger *Trigger) error
}
