package model

import (
	"errors"
	"time"
)

type Status int

type TopicInfo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`

	SubscribeTarget SubscriberType `json:"subscribe_target"`
	TriggerId       int64          `json:"trigger_id"`
	// todo cms
}

type TopicTemplate struct {
	ID         int64     `json:"id"`
	TopicId    int64     `json:"topic_id"`
	Channel    Channel   `json:"channel"`
	Subject    string    `json:"subject"`
	Content    string    `json:"content"`
	Status     Status    `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type AggregateTopic struct {
	TopicInfo TopicInfo        `json:"topicInfo"`
	Templates []*TopicTemplate `json:"templates"`
}

func (t *AggregateTopic) selectTemplate(channel Channel) (*TopicTemplate, error) {
	for _, template := range t.Templates {
		if template.Channel == channel {
			return template, nil
		}
	}
	return &TopicTemplate{}, errors.New("template not found")
}
