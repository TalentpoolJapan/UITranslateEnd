package repo

import (
	"log"
	"uitranslate/domain/notification/model"
	"uitranslate/infrastructure"
	"xorm.io/xorm"
)

var (
	Repo = NewNotificationRepository(infrastructure.MysqlDB)
)

type NotificationRepository struct {
	DB *xorm.Engine
}

func NewNotificationRepository(db *xorm.Engine) *NotificationRepository {
	return &NotificationRepository{
		DB: db,
	}
}

// region topic info

func (r *NotificationRepository) GetTopicInfoById(topicId int64) (*model.TopicInfo, error) {
	var topicInfoPO TopicInfoPO
	has, err := r.DB.Table(TopicInfoTableName).ID(topicId).Get(&topicInfoPO)
	if err != nil {
		log.Printf("Error fetching topic info: %v", err)
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return topicInfoPO.ConvertToEntity(), nil
}

func (r *NotificationRepository) ListTopicInfo() ([]*model.TopicInfo, error) {
	var topicInfoPOs []*TopicInfoPO
	err := r.DB.Table(TopicInfoTableName).Find(&topicInfoPOs)
	if err != nil {
		log.Printf("Error fetching topic info: %v", err)
		return nil, err
	}
	var topicInfos []*model.TopicInfo
	for _, topicInfoPO := range topicInfoPOs {
		topicInfos = append(topicInfos, topicInfoPO.ConvertToEntity())
	}
	return topicInfos, nil
}

func (r *NotificationRepository) SaveTopicInfo(topicInfo model.TopicInfo) error {
	topicInfoPO := ConvertTopicInfoPO(topicInfo)
	_, err := r.DB.Table(TopicInfoTableName).Insert(topicInfoPO)
	if err != nil {
		log.Printf("Error saving topic info: %v", err)
		return err
	}
	return nil
}

func (r *NotificationRepository) UpdateTopicInfo(topicInfo model.TopicInfo) error {
	topicInfoPO := ConvertTopicInfoPO(topicInfo)
	_, err := r.DB.Table(TopicInfoTableName).ID(topicInfo.ID).Update(topicInfoPO)
	if err != nil {
		log.Printf("Error updating topic info: %v", err)
		return err
	}
	return nil
}

// endregion

// region topic template

func (r *NotificationRepository) GetTopicTemplateById(templateId int64) (*model.TopicTemplate, error) {
	var topicTemplatePO TopicTemplatePO
	has, err := r.DB.Table(TopicTemplateTableName).ID(templateId).Get(&topicTemplatePO)
	if err != nil {
		log.Printf("Error fetching topic template: %v", err)
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return topicTemplatePO.ConvertToEntity(), nil
}

func (r *NotificationRepository) ListTopicTemplate() ([]*model.TopicTemplate, error) {
	var topicTemplatePOs []*TopicTemplatePO
	err := r.DB.Table(TopicTemplateTableName).Find(&topicTemplatePOs)
	if err != nil {
		log.Printf("Error fetching topic templates: %v", err)
		return nil, err
	}
	var topicTemplates []*model.TopicTemplate
	for _, topicTemplatePO := range topicTemplatePOs {
		topicTemplates = append(topicTemplates, topicTemplatePO.ConvertToEntity())
	}
	return topicTemplates, nil
}

func (r *NotificationRepository) ListTopicTemplateByTopicId(topicId int64) ([]*model.TopicTemplate, error) {
	var topicTemplatePOs []*TopicTemplatePO
	err := r.DB.Table(TopicTemplateTableName).Where("topic_id = ?", topicId).Find(&topicTemplatePOs)
	if err != nil {
		log.Printf("Error fetching topic templates: %v", err)
		return nil, err
	}
	var topicTemplates []*model.TopicTemplate
	for _, topicTemplatePO := range topicTemplatePOs {
		topicTemplates = append(topicTemplates, topicTemplatePO.ConvertToEntity())
	}
	return topicTemplates, nil
}

func (r *NotificationRepository) SaveTopicTemplate(template model.TopicTemplate) error {
	topicTemplatePO := ConvertTopicTemplatePO(template)
	_, err := r.DB.Table(TopicTemplateTableName).Insert(topicTemplatePO)
	if err != nil {
		log.Printf("Error saving topic template: %v", err)
		return err
	}
	return nil
}

func (r *NotificationRepository) UpdateTopicTemplate(template model.TopicTemplate) error {
	topicTemplatePO := ConvertTopicTemplatePO(template)
	_, err := r.DB.Table(TopicTemplateTableName).ID(template.ID).Update(topicTemplatePO)
	if err != nil {
		log.Printf("Error updating topic template: %v", err)
		return err
	}
	return nil
}

// endregion

// region trigger

func (r *NotificationRepository) GetTriggerById(triggerId int64) (*model.Trigger, error) {
	var triggerPO TriggerPO
	has, err := r.DB.Table(TriggerTableName).ID(triggerId).Get(&triggerPO)
	if err != nil {
		log.Printf("Error fetching trigger: %v", err)
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return triggerPO.ConvertToEntity(), nil
}

func (r *NotificationRepository) ListTrigger() ([]*model.Trigger, error) {
	var triggerPOs []*TriggerPO
	err := r.DB.Table(TriggerTableName).Find(&triggerPOs)
	if err != nil {
		log.Printf("Error fetching triggers: %v", err)
		return nil, err
	}
	var triggers []*model.Trigger
	for _, triggerPO := range triggerPOs {
		triggers = append(triggers, triggerPO.ConvertToEntity())
	}
	return triggers, nil
}

func (r *NotificationRepository) ListTriggerByTopicId(topicId int64) ([]*model.Trigger, error) {
	var triggerPOs []*TriggerPO
	err := r.DB.Table(TriggerTableName).Where("topic_id = ?", topicId).Find(&triggerPOs)
	if err != nil {
		log.Printf("Error fetching triggers: %v", err)
		return nil, err
	}
	var triggers []*model.Trigger
	for _, triggerPO := range triggerPOs {
		triggers = append(triggers, triggerPO.ConvertToEntity())
	}
	return triggers, nil
}

func (r *NotificationRepository) SaveTrigger(trigger model.Trigger) error {
	triggerPO := ConvertTriggerPO(trigger)
	_, err := r.DB.Table(TriggerTableName).Insert(triggerPO)
	if err != nil {
		log.Printf("Error saving trigger: %v", err)
		return err
	}
	return nil
}

func (r *NotificationRepository) UpdateTrigger(trigger model.Trigger) error {
	triggerPO := ConvertTriggerPO(trigger)
	_, err := r.DB.Table(TriggerTableName).ID(trigger.ID).Update(triggerPO)
	if err != nil {
		log.Printf("Error updating trigger: %v", err)
		return err
	}
	return nil
}

// endregion

// region subscribe

func (r *NotificationRepository) ListSubscribeTopicMappingBySubscriber(subscriber *model.Subscriber) ([]*model.SubscribeTopicMapping, error) {
	var subscribeTopicMappingPOs []*SubscribeTopicMappingPO
	err := r.DB.Table(SubscriberTopicMappingTableName).Where("subscriber_uuid = ?", subscriber.Uuid).Find(&subscribeTopicMappingPOs)
	if err != nil {
		return nil, err
	}
	var subscribeTopicMappings []*model.SubscribeTopicMapping
	for _, subscribeTopicMappingPO := range subscribeTopicMappingPOs {
		subscribeTopicMappings = append(subscribeTopicMappings, subscribeTopicMappingPO.ConvertToEntity())
	}
	return subscribeTopicMappings, nil
}

func (r *NotificationRepository) SaveSubscribeTopicMapping(mapping []*model.SubscribeTopicMapping) error {
	session := r.DB.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return err
	}

	for _, mapEntity := range mapping {
		mapPO := ConvertSubscribeTopicMappingPO(*mapEntity)
		_, err := session.Table(SubscriberTopicMappingTableName).Insert(mapPO)
		if err != nil {
			_ = session.Rollback()
			return err
		}
	}

	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *NotificationRepository) RemoveSubscribeTopicMappingBySubscriberId(subscriberUuid string) error {
	_, err := r.DB.Table(SubscriberTopicMappingTableName).Where("subscriber_uuid = ?", subscriberUuid).Delete(&SubscribeTopicMappingPO{})
	if err != nil {
		return err
	}

	return nil
}

// endregion
