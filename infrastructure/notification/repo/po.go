package repo

import (
	"time"
	"uitranslate/domain/notification"
	"uitranslate/domain/notification/model"
	"uitranslate/domain/notification/subscribe"
	"uitranslate/domain/notification/topic"
)

const (
	TopicInfoTableName              = "nf_topic_info"
	TopicTemplateTableName          = "nf_topic_template"
	TriggerTableName                = "nf_trigger"
	SubscriberTopicMappingTableName = "nf_subscriber_topic_mapping"
)

type TopicInfoPO struct {
	Id              int64     `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Status          int       `json:"status"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
	SubscribeTarget string    `json:"subscribe_target"`
	TriggerId       int64     `json:"trigger_id"`
}

type TopicTemplatePO struct {
	Id         int64                `json:"id"`
	TopicId    int64                `json:"topic"`
	Name       string               `json:"name"`
	Channel    notification.Channel `json:"channel"`
	Subject    string               `json:"subject"`
	Content    string               `json:"content"`
	Status     int                  `json:"status"`
	CreateTime time.Time            `json:"create_time"`
	UpdateTime time.Time            `json:"update_time"`
}

type TriggerPO struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Immediate  int       `json:"immediate"`
	Frequency  string    `json:"frequency"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type SubscribeTopicMappingPO struct {
	Id             int64                    `json:"id"`
	SubscriberType subscribe.SubscriberType `json:"subscriber_type"`
	SubscriberUuid string                   `json:"subscriber_uuid"`
	TopicId        int64                    `json:"topic_id"`
}

func ConvertTopicInfoPO(entity topic.TopicInfo) *TopicInfoPO {
	return &TopicInfoPO{
		Id:              entity.ID,
		Title:           entity.Title,
		Description:     entity.Description,
		Status:          int(entity.Status),
		CreateTime:      entity.CreateTime,
		UpdateTime:      entity.UpdateTime,
		SubscribeTarget: entity.SubscribeTarget,
		TriggerId:       entity.TriggerId,
	}
}

func (po *TopicInfoPO) ConvertToEntity() *topic.TopicInfo {
	return &topic.TopicInfo{
		ID:              po.Id,
		Title:           po.Title,
		Description:     po.Description,
		Status:          topic.Status(po.Status),
		CreateTime:      po.CreateTime,
		UpdateTime:      po.UpdateTime,
		SubscribeTarget: po.SubscribeTarget,
		TriggerId:       po.TriggerId,
	}
}

func ConvertTopicTemplatePO(entity topic.TopicTemplate) *TopicTemplatePO {
	return &TopicTemplatePO{
		Id:         entity.ID,
		TopicId:    entity.TopicId,
		Channel:    entity.Channel,
		Name:       entity.Name,
		Subject:    entity.Subject,
		Content:    entity.Content,
		Status:     int(entity.Status),
		CreateTime: entity.CreateTime,
		UpdateTime: entity.UpdateTime,
	}
}

func (po *TopicTemplatePO) ConvertToEntity() *topic.TopicTemplate {
	return &topic.TopicTemplate{
		ID:         po.Id,
		TopicId:    po.TopicId,
		Channel:    po.Channel,
		Name:       po.Name,
		Subject:    po.Subject,
		Content:    po.Content,
		Status:     topic.Status(po.Status),
		CreateTime: po.CreateTime,
		UpdateTime: po.UpdateTime,
	}
}

func ConvertTriggerPO(entity model.Trigger) *TriggerPO {
	return &TriggerPO{
		Id:         entity.ID,
		Name:       entity.Name,
		Immediate:  boolToInt(entity.Immediate),
		Frequency:  entity.Frequency,
		CreateTime: entity.CreateTime,
		UpdateTime: entity.UpdateTime,
	}
}

func (po *TriggerPO) ConvertToEntity() *model.Trigger {
	return &model.Trigger{
		ID:         po.Id,
		Name:       po.Name,
		Immediate:  intToBool(po.Immediate),
		Frequency:  po.Frequency,
		CreateTime: po.CreateTime,
		UpdateTime: po.UpdateTime,
	}
}

func ConvertSubscribeTopicMappingPO(entity *subscribe.SubscribeTopicMapping) *SubscribeTopicMappingPO {
	return &SubscribeTopicMappingPO{
		Id:             entity.ID,
		SubscriberType: entity.SubscriberType,
		SubscriberUuid: entity.SubscriberUuid,
		TopicId:        entity.TopicId,
	}
}

func (po *SubscribeTopicMappingPO) ConvertToEntity() *subscribe.SubscribeTopicMapping {
	return &subscribe.SubscribeTopicMapping{
		ID:             po.Id,
		SubscriberType: po.SubscriberType,
		SubscriberUuid: po.SubscriberUuid,
		TopicId:        po.TopicId,
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func intToBool(i int) bool {
	return i == 1
}
