package impl

import (
	"fmt"
	"uitranslate/app/notification"
	"uitranslate/domain/notification/gateway"
	"uitranslate/domain/notification/model"
	inf "uitranslate/infrastructure/notification"
)

var (
	TopicAppServSingleton = NewTopicAppServImpl()
)

type TopicAppServImpl struct {
	gateway gateway.Gateway
}

func NewTopicAppServImpl() notification.TopicAppServ {
	return &TopicAppServImpl{
		gateway: inf.GatewaySingleton,
	}
}

func (t *TopicAppServImpl) GetTopicInfoById(qry notification.TopicInfoByIdQry) (*notification.TopicInfoResp, error) {
	topicInfo, err := t.gateway.GetTopicInfoById(qry.ID)
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
	topicInfos, err := t.gateway.ListTopicInfo()
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
	return t.gateway.SaveTopicInfo(&model.TopicInfo{
		Title:           cmd.Title,
		Description:     cmd.Description,
		Status:          model.Status(cmd.Status),
		SubscribeTarget: cmd.SubscribeTarget,
		TriggerId:       cmd.TriggerId,
	})
}

func (t *TopicAppServImpl) UpdateTopicInfo(cmd notification.TopicInfoUpdateCmd) error {
	return t.gateway.UpdateTopicInfo(&model.TopicInfo{
		ID:              cmd.ID,
		Title:           cmd.Title,
		Description:     cmd.Description,
		Status:          model.Status(cmd.Status),
		SubscribeTarget: cmd.SubscribeTarget,
		TriggerId:       cmd.TriggerId,
	})
}

func (t *TopicAppServImpl) GetTopicTemplateById(qry notification.TopicTemplateByIdQry) (*notification.TopicTemplateResp, error) {
	topicTemplate, err := t.gateway.GetTopicTemplateById(qry.ID)
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
		Subject:    topicTemplate.Subject,
		Content:    topicTemplate.Content,
		Status:     int(topicTemplate.Status),
		CreateTime: topicTemplate.CreateTime,
		UpdateTime: topicTemplate.UpdateTime,
	}, nil
}

func (t *TopicAppServImpl) ListTopicTemplateByTopicId(qry notification.TopicTemplateByTopicIdQuery) ([]*notification.TopicTemplateResp, error) {
	topicTemplates, err := t.gateway.ListTopicTemplateByTopicId(qry.TopicId)
	if err != nil {
		return nil, err
	}
	var resp []*notification.TopicTemplateResp
	for _, topicTemplate := range topicTemplates {
		resp = append(resp, &notification.TopicTemplateResp{
			ID:         topicTemplate.ID,
			TopicId:    topicTemplate.TopicId,
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
	return t.gateway.SaveTopicTemplate(&model.TopicTemplate{
		TopicId: cmd.TopicId,
		Channel: model.Channel(cmd.Channel),
		Subject: cmd.Subject,
		Content: cmd.Content,
		Status:  model.Status(cmd.Status),
	})
}

func (t *TopicAppServImpl) UpdateTopicTemplate(cmd notification.TopicTemplateUpdateCmd) error {
	return t.gateway.UpdateTopicTemplate(&model.TopicTemplate{
		ID:      cmd.ID,
		TopicId: cmd.TopicId,
		Channel: model.Channel(cmd.Channel),
		Subject: cmd.Subject,
		Content: cmd.Content,
		Status:  model.Status(cmd.Status),
	})
}
