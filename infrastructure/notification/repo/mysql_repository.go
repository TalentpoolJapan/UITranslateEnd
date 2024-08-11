package repo

import (
	"log"
	"uitranslate/domain/notification/subscriber"
	"uitranslate/domain/notification/topic"
	"uitranslate/domain/notification/trigger"
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

func (r *NotificationRepository) GetTopicInfoById(topicId int64) (*topic.BasicInfo, error) {
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

func (r *NotificationRepository) ListTopicInfo() ([]*topic.BasicInfo, error) {
	var topicInfoPOs []*TopicInfoPO
	err := r.DB.Table(TopicInfoTableName).Find(&topicInfoPOs)
	if err != nil {
		log.Printf("Error fetching topic info: %v", err)
		return nil, err
	}
	var topicInfos []*topic.BasicInfo
	for _, topicInfoPO := range topicInfoPOs {
		topicInfos = append(topicInfos, topicInfoPO.ConvertToEntity())
	}
	return topicInfos, nil
}

func (r *NotificationRepository) SaveTopicInfo(topicInfo topic.BasicInfo) error {
	topicInfoPO := ConvertTopicInfoPO(topicInfo)
	_, err := r.DB.Table(TopicInfoTableName).Insert(topicInfoPO)
	if err != nil {
		log.Printf("Error saving topic info: %v", err)
		return err
	}
	return nil
}

func (r *NotificationRepository) UpdateTopicInfo(topicInfo topic.BasicInfo) error {
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

func (r *NotificationRepository) GetTopicTemplateById(templateId int64) (*topic.Template, error) {
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

func (r *NotificationRepository) ListTopicTemplate() ([]*topic.Template, error) {
	var topicTemplatePOs []*TopicTemplatePO
	err := r.DB.Table(TopicTemplateTableName).Find(&topicTemplatePOs)
	if err != nil {
		log.Printf("Error fetching topic templates: %v", err)
		return nil, err
	}
	var topicTemplates []*topic.Template
	for _, topicTemplatePO := range topicTemplatePOs {
		topicTemplates = append(topicTemplates, topicTemplatePO.ConvertToEntity())
	}
	return topicTemplates, nil
}

func (r *NotificationRepository) ListTopicTemplateByTopicId(topicId int64) ([]*topic.Template, error) {
	var topicTemplatePOs []*TopicTemplatePO
	err := r.DB.Table(TopicTemplateTableName).Where("topic_id = ?", topicId).Find(&topicTemplatePOs)
	if err != nil {
		log.Printf("Error fetching topic templates: %v", err)
		return nil, err
	}
	var topicTemplates []*topic.Template
	for _, topicTemplatePO := range topicTemplatePOs {
		topicTemplates = append(topicTemplates, topicTemplatePO.ConvertToEntity())
	}
	return topicTemplates, nil
}

func (r *NotificationRepository) SaveTopicTemplate(template topic.Template) error {
	topicTemplatePO := ConvertTopicTemplatePO(template)
	_, err := r.DB.Table(TopicTemplateTableName).Insert(topicTemplatePO)
	if err != nil {
		log.Printf("Error saving topic template: %v", err)
		return err
	}
	return nil
}

func (r *NotificationRepository) UpdateTopicTemplate(template topic.Template) error {
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

func (r *NotificationRepository) GetTriggerById(triggerId int64) (*trigger.Trigger, error) {
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

func (r *NotificationRepository) ListTrigger() ([]*trigger.Trigger, error) {
	var triggerPOs []*TriggerPO
	err := r.DB.Table(TriggerTableName).Find(&triggerPOs)
	if err != nil {
		log.Printf("Error fetching triggers: %v", err)
		return nil, err
	}
	var triggers []*trigger.Trigger
	for _, triggerPO := range triggerPOs {
		triggers = append(triggers, triggerPO.ConvertToEntity())
	}
	return triggers, nil
}

func (r *NotificationRepository) ListTriggerByTopicId(topicId int64) ([]*trigger.Trigger, error) {
	var triggerPOs []*TriggerPO
	err := r.DB.Table(TriggerTableName).Where("topic_id = ?", topicId).Find(&triggerPOs)
	if err != nil {
		log.Printf("Error fetching triggers: %v", err)
		return nil, err
	}
	var triggers []*trigger.Trigger
	for _, triggerPO := range triggerPOs {
		triggers = append(triggers, triggerPO.ConvertToEntity())
	}
	return triggers, nil
}

func (r *NotificationRepository) SaveTrigger(trigger trigger.Trigger) error {
	triggerPO := ConvertTriggerPO(trigger)
	_, err := r.DB.Table(TriggerTableName).Insert(triggerPO)
	if err != nil {
		log.Printf("Error saving trigger: %v", err)
		return err
	}
	return nil
}

func (r *NotificationRepository) UpdateTrigger(trigger trigger.Trigger) error {
	triggerPO := ConvertTriggerPO(trigger)
	_, err := r.DB.Table(TriggerTableName).ID(trigger.ID).Update(triggerPO)
	if err != nil {
		log.Printf("Error updating trigger: %v", err)
		return err
	}
	return nil
}

// endregion

// region subscriber

func (r *NotificationRepository) ListSubscribedTopicIdsBySubscriber(userType subscriber.Type, userUuid string) ([]int64, error) {
	var subscribeTopicMappingPOs []*SubscribeTopicMappingPO
	err := r.DB.Table(SubscriberTopicMappingTableName).Where("subscriber_uuid = ?", userUuid).And("subscriber_type", userType).Find(&subscribeTopicMappingPOs)
	if err != nil {
		return nil, err
	}
	topicIds := make([]int64, len(subscribeTopicMappingPOs))
	for _, subscribeTopicMappingPO := range subscribeTopicMappingPOs {
		topicIds = append(topicIds, subscribeTopicMappingPO.TopicId)
	}
	return topicIds, nil
}

func (r *NotificationRepository) SaveSubscriberTopic(subscriber *subscriber.Subscriber) error {
	// start transaction
	session := r.DB.NewSession()
	err := session.Begin()
	if err != nil {
		return err
	}
	// delete mapping by subscriber
	delSql := "delete from ? where subscriber_uuid = ? and subscriber_type = ?"
	_, err = session.Exec(delSql, SubscriberTopicMappingTableName, subscriber.Uuid, subscriber.Type)
	if err != nil {
		return err
	}

	// save batch topic 可批量优化
	for _, topicId := range subscriber.TopicIds {
		po := SubscribeTopicMappingPO{
			SubscriberType: subscriber.Type,
			SubscriberUuid: subscriber.Uuid,
			TopicId:        topicId,
		}
		_, err := session.Table(SubscriberTopicMappingTableName).Insert(po)
		if err != nil {
			_ = session.Rollback()
			return err
		}
	}

	// commit
	err = session.Commit()
	return err
}

// endregion
