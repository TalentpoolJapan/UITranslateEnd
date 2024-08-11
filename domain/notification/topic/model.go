package topic

import (
	"errors"
	"time"
	"uitranslate/domain/notification"
)

type Status int

type BasicInfo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`

	SubscribeTarget string `json:"subscribe_target"`
	TriggerId       int64  `json:"trigger_id"`
}

type Template struct {
	ID         int64                `json:"id"`
	TopicId    int64                `json:"topic_id"`
	Name       string               `json:"name"`
	Channel    notification.Channel `json:"channel"`
	Subject    string               `json:"subject"`
	Content    string               `json:"content"`
	Status     Status               `json:"status"`
	CreateTime time.Time            `json:"create_time"`
	UpdateTime time.Time            `json:"update_time"`
}

type AggregateTopic struct {
	TopicInfo BasicInfo   `json:"topicInfo"`
	Templates []*Template `json:"templates"`
}

func (t *AggregateTopic) SelectTemplate(channel notification.Channel) (*Template, error) {
	for _, template := range t.Templates {
		if template.Channel == channel {
			return template, nil
		}
	}
	return &Template{}, errors.New("template not found")
}
