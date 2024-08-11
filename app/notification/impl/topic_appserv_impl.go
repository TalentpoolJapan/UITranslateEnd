package impl

import (
	"fmt"
	"uitranslate/app/notification"
	notification2 "uitranslate/domain/notification"
	"uitranslate/domain/notification/topic"
)

type TopicAppServImpl struct {
	topicGateway topic.Repository
}

func NewTopicAppServImpl(topicGateway topic.Repository) notification.TopicAppServ {
	return &TopicAppServImpl{
		topicGateway: topicGateway,
	}
}

func (t *TopicAppServImpl) GetTopicInfoById(qry notification.TopicInfoByIdQry) (*notification.TopicInfoResp, error) {
	topicInfo, err := t.topicGateway.GetTopicInfoById(qry.ID)
	if err != nil {
		return nil, err
	}
	return &notification.TopicInfoResp{
		ID:              topicInfo.ID,
		Title:           topicInfo.Title,
		Description:     topicInfo.Description,
		Status:          int(topicInfo.Status),
		CreateTime:      topicInfo.CreateTime,
		UpdateTime:      topicInfo.UpdateTime,
		SubscribeTarget: topicInfo.SubscribeTarget,
		TriggerId:       topicInfo.TriggerId,
	}, nil
}

func (t *TopicAppServImpl) ListTopicInfo() ([]*notification.TopicInfoResp, error) {
	topicInfos, err := t.topicGateway.ListTopicInfo()
	if err != nil {
		return nil, err
	}
	var resp []*notification.TopicInfoResp
	for _, topicInfo := range topicInfos {
		resp = append(resp, &notification.TopicInfoResp{
			ID:              topicInfo.ID,
			Title:           topicInfo.Title,
			Description:     topicInfo.Description,
			Status:          int(topicInfo.Status),
			CreateTime:      topicInfo.CreateTime,
			UpdateTime:      topicInfo.UpdateTime,
			SubscribeTarget: topicInfo.SubscribeTarget,
			TriggerId:       topicInfo.TriggerId,
		})
	}
	return resp, nil
}

func (t *TopicAppServImpl) AddTopicInfo(cmd notification.TopicInfoAddCmd) error {
	return t.topicGateway.SaveTopicInfo(&topic.BasicInfo{
		Title:           cmd.Title,
		Description:     cmd.Description,
		Status:          topic.Status(cmd.Status),
		SubscribeTarget: cmd.SubscribeTarget,
		TriggerId:       cmd.TriggerId,
	})
}

func (t *TopicAppServImpl) UpdateTopicInfo(cmd notification.TopicInfoUpdateCmd) error {
	return t.topicGateway.UpdateTopicInfo(&topic.BasicInfo{
		ID:              cmd.ID,
		Title:           cmd.Title,
		Description:     cmd.Description,
		Status:          topic.Status(cmd.Status),
		SubscribeTarget: cmd.SubscribeTarget,
		TriggerId:       cmd.TriggerId,
	})
}

func (t *TopicAppServImpl) GetTopicTemplateById(qry notification.TopicTemplateByIdQry) (*notification.TopicTemplateResp, error) {
	topicTemplate, err := t.topicGateway.GetTopicTemplateById(qry.ID)
	if err != nil {
		return nil, err
	}
	if topicTemplate == nil {
		return nil, fmt.Errorf("TopicTemplate with Id %d not found", qry.ID)
	}
	return &notification.TopicTemplateResp{
		ID:         topicTemplate.ID,
		TopicId:    topicTemplate.TopicId,
		Channel:    string(topicTemplate.Channel),
		Name:       topicTemplate.Name,
		Subject:    topicTemplate.Subject,
		Content:    topicTemplate.Content,
		Status:     int(topicTemplate.Status),
		CreateTime: topicTemplate.CreateTime,
		UpdateTime: topicTemplate.UpdateTime,
	}, nil
}

func (t *TopicAppServImpl) ListTopicTemplateByTopicId(qry notification.TopicTemplateByTopicIdQuery) ([]*notification.TopicTemplateResp, error) {
	topicTemplates, err := t.topicGateway.ListTopicTemplateByTopicId(qry.TopicId)
	if err != nil {
		return nil, err
	}
	var resp []*notification.TopicTemplateResp
	for _, topicTemplate := range topicTemplates {
		resp = append(resp, &notification.TopicTemplateResp{
			ID:         topicTemplate.ID,
			TopicId:    topicTemplate.TopicId,
			Name:       topicTemplate.Name,
			Channel:    string(topicTemplate.Channel),
			Subject:    topicTemplate.Subject,
			Content:    topicTemplate.Content,
			Status:     int(topicTemplate.Status),
			CreateTime: topicTemplate.CreateTime,
			UpdateTime: topicTemplate.UpdateTime,
		})
	}
	return resp, nil
}

func (t *TopicAppServImpl) AddTopicTemplate(cmd notification.TopicTemplateAddCmd) error {
	return t.topicGateway.SaveTopicTemplate(&topic.Template{
		TopicId: cmd.TopicId,
		Name:    cmd.Name,
		Channel: notification2.Channel(cmd.Channel),
		Subject: cmd.Subject,
		Content: cmd.Content,
		Status:  topic.Status(cmd.Status),
	})
}

func (t *TopicAppServImpl) UpdateTopicTemplate(cmd notification.TopicTemplateUpdateCmd) error {
	return t.topicGateway.UpdateTopicTemplate(&topic.Template{
		ID:      cmd.ID,
		TopicId: cmd.TopicId,
		Name:    cmd.Name,
		Channel: notification2.Channel(cmd.Channel),
		Subject: cmd.Subject,
		Content: cmd.Content,
		Status:  topic.Status(cmd.Status),
	})
}
