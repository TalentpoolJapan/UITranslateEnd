package notification

import (
	"time"
)

type TopicInfoResp struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`

	SubscribeTarget string `json:"subscribe_target"`
	TriggerId       int64  `json:"trigger_id"`
}

type TopicTemplateResp struct {
	ID         int64     `json:"id"`
	TopicId    int64     `json:"topic_id"`
	Channel    string    `json:"channel"`
	Subject    string    `json:"subject"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type TopicInfoByIdQry struct {
	ID int64
}

type TopicInfoAddCmd struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Status          int    `json:"status"`
	SubscribeTarget string `json:"subscribe_target"`
	TriggerId       int64  `json:"trigger_id"`
}

type TopicInfoUpdateCmd struct {
	ID              int64  `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Status          int    `json:"status"`
	SubscribeTarget string `json:"subscribe_target"`
	TriggerId       int64  `json:"trigger_id"`
}

type TopicTemplateByIdQry struct {
	ID int64
}

type TopicTemplateByTopicIdQuery struct {
	TopicId int64
}

type TopicTemplateAddCmd struct {
	TopicId int64  `json:"topic_id"`
	Channel string `json:"channel"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}

type TopicTemplateUpdateCmd struct {
	ID      int64  `json:"id"`
	TopicId int64  `json:"topic_id"`
	Channel string `json:"channel"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}

type TopicAppServ interface {
	GetTopicInfoById(qry TopicInfoByIdQry) (*TopicInfoResp, error)
	ListTopicInfo() ([]*TopicInfoResp, error)
	AddTopicInfo(cmd TopicInfoAddCmd) error
	UpdateTopicInfo(cmd TopicInfoUpdateCmd) error

	GetTopicTemplateById(qry TopicTemplateByIdQry) (*TopicTemplateResp, error)
	ListTopicTemplateByTopicId(qry TopicTemplateByTopicIdQuery) ([]*TopicTemplateResp, error)
	AddTopicTemplate(cmd TopicTemplateAddCmd) error
	UpdateTopicTemplate(cmd TopicTemplateUpdateCmd) error
}
