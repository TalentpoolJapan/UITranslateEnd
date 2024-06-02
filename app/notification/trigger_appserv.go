package notification

import "time"

type TriggerResp struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Immediate  bool      `json:"immediate"`
	Frequency  string    `json:"frequency"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type TriggerByIdQry struct {
	ID int64 `json:"id"`
}

type TriggerAddCmd struct {
	Name      string `json:"name"`
	Immediate bool   `json:"immediate"`
	Frequency string `json:"frequency"`
}

type TriggerUpdateCmd struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Immediate bool   `json:"immediate"`
	Frequency string `json:"frequency"`
}

type TriggerAppServ interface {
	GetTriggerById(qry *TriggerByIdQry) (*TriggerResp, error)
	ListTrigger() ([]*TriggerResp, error)
	AddTrigger(cmd *TriggerAddCmd) error
	UpdateTrigger(cmd *TriggerUpdateCmd) error
}
