package notification

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
	Topic     topic.BasicInfo `json:"topic"`
}
