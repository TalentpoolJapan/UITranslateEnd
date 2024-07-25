package gateway

import "uitranslate/domain/notification/topic"

type TopicGateway interface {
	GetTopicInfoById(topicId int64) (*topic.TopicInfo, error)
	ListTopicInfo() ([]*topic.TopicInfo, error)
	SaveTopicInfo(topicInfo *topic.TopicInfo) error
	UpdateTopicInfo(topicInfo *topic.TopicInfo) error

	GetTopicTemplateById(templateId int64) (*topic.TopicTemplate, error)
	ListTopicTemplateByTopicId(topicId int64) ([]*topic.TopicTemplate, error)
	SaveTopicTemplate(topicTemplate *topic.TopicTemplate) error
	UpdateTopicTemplate(topicTemplate *topic.TopicTemplate) error
}
