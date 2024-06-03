package repo

import (
	"time"
	"uitranslate/domain/notification/model"
)

const (
	TopicInfoTableName              = "nf_topic_info"
	TopicTemplateTableName          = "nf_topic_template"
	TriggerTableName                = "nf_trigger"
	SubscriberTopicMappingTableName = "nf_subscriber_topic_mapping"
)

type TopicInfoPO struct {
	Id              int64                `json:"id"`
	Title           string               `json:"title"`
	Description     string               `json:"description"`
	Status          int                  `json:"status"`
	CreateTime      time.Time            `json:"create_time"`
	UpdateTime      time.Time            `json:"update_time"`
	SubscribeTarget model.SubscriberType `json:"subscribe_target"`
	TriggerId       int64                `json:"trigger_id"`
}

type TopicTemplatePO struct {
	Id         int64         `json:"id"`
	TopicId    int64         `json:"topic"`
	Channel    model.Channel `json:"channel"`
	Subject    string        `json:"subject"`
	Content    string        `json:"content"`
	Status     int           `json:"status"`
	CreateTime time.Time     `json:"create_time"`
	UpdateTime time.Time     `json:"update_time"`
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
	SubscriberType model.SubscriberType `json:"subscriber_type"`
	SubscriberUuid int64                `json:"subscriber_uuid"`
	TopicId        int64                `json:"topic_id"`
}

func ConvertTopicInfoPO(entity model.TopicInfo) *TopicInfoPO {
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

func (po *TopicInfoPO) ConvertToEntity() *model.TopicInfo {
	return &model.TopicInfo{
		ID:              po.Id,
		Title:           po.Title,
		Description:     po.Description,
		Status:          model.Status(po.Status),
		CreateTime:      po.CreateTime,
		UpdateTime:      po.UpdateTime,
		SubscribeTarget: po.SubscribeTarget,
		TriggerId:       po.TriggerId,
	}
}

func ConvertTopicTemplatePO(entity model.TopicTemplate) *TopicTemplatePO {
	return &TopicTemplatePO{
		Id:         entity.ID,
		TopicId:    entity.TopicId,
		Channel:    entity.Channel,
		Subject:    entity.Subject,
		Content:    entity.Content,
		Status:     int(entity.Status),
		CreateTime: entity.CreateTime,
		UpdateTime: entity.UpdateTime,
	}
}

func (po *TopicTemplatePO) ConvertToEntity() *model.TopicTemplate {
	return &model.TopicTemplate{
		ID:         po.Id,
		TopicId:    po.TopicId,
		Channel:    po.Channel,
		Subject:    po.Subject,
		Content:    po.Content,
		Status:     model.Status(po.Status),
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

func ConvertSubscribeTopicMappingPO(entity model.SubscribeTopicMapping) *SubscribeTopicMappingPO {
	return &SubscribeTopicMappingPO{
		SubscriberType: entity.SubscriberType,
		SubscriberUuid: entity.SubscriberUuid,
		TopicId:        entity.TopicId,
	}
}

func (po *SubscribeTopicMappingPO) ConvertToEntity() *model.SubscribeTopicMapping {
	return &model.SubscribeTopicMapping{
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
