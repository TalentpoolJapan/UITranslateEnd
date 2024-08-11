package trigger

import "time"

type Trigger struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Immediate  bool      `json:"immediate"`
	Frequency  string    `json:"frequency"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
