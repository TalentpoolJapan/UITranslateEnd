package model

import "uitranslate/domain/notification/topic"

type EventType string

const (
	JobPosted              EventType = "JobPosted"
	JobCancelled           EventType = "JobCancelled"
	CompanyUpdated         EventType = "CompanyUpdated"
	JobApplicationReceived EventType = "JobApplicationReceived"
)

type EventTypeSetting struct {
	EventType EventType       `json:"event_type"`
	Topic     topic.TopicInfo `json:"topic"`
}
