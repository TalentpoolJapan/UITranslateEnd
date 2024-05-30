package notification

import "time"

type Status int

type Topic struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`

	SubscribeTarget string  `json:"subscribe_target"`
	Trigger         Trigger `json:"trigger"`
	// todo cms
}

type Trigger struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Immediate bool   `json:"immediate"`
	Frequency string `json:"frequency"`
}

type Template struct {
	ID      int64  `json:"id"`
	Topic   Topic  `json:"topic"`
	Channel string `json:"channel"`
	Content string `json:"content"`
	Status  Status `json:"status"`
}
