package impl

import (
	"fmt"
	"uitranslate/app/notification"
	"uitranslate/domain/notification/gateway"
	"uitranslate/domain/notification/model"
)

type TriggerAppServImpl struct {
	gateway gateway.Gateway
}

func NewTriggerAppServImpl(g gateway.Gateway) notification.TriggerAppServ {
	return &TriggerAppServImpl{
		gateway: g,
	}
}

func (t *TriggerAppServImpl) GetTriggerById(qry *notification.TriggerByIdQry) (*notification.TriggerResp, error) {
	trigger, err := t.gateway.GetTriggerById(qry.ID)
	if err != nil {
		return nil, err
	}
	if trigger == nil {
		return nil, fmt.Errorf("Trigger with Id %d not found", qry.ID)
	}
	return &notification.TriggerResp{
		ID:         trigger.ID,
		Name:       trigger.Name,
		Immediate:  trigger.Immediate,
		Frequency:  trigger.Frequency,
		CreateTime: trigger.CreateTime,
		UpdateTime: trigger.UpdateTime,
	}, nil
}

func (t *TriggerAppServImpl) ListTrigger() ([]*notification.TriggerResp, error) {
	triggers, err := t.gateway.ListTrigger()
	if err != nil {
		return nil, err
	}
	var resp []*notification.TriggerResp
	for _, trigger := range triggers {
		resp = append(resp, &notification.TriggerResp{
			ID:         trigger.ID,
			Name:       trigger.Name,
			Immediate:  trigger.Immediate,
			Frequency:  trigger.Frequency,
			CreateTime: trigger.CreateTime,
			UpdateTime: trigger.UpdateTime,
		})
	}
	return resp, nil
}

func (t *TriggerAppServImpl) AddTrigger(cmd *notification.TriggerAddCmd) error {
	return t.gateway.SaveTrigger(&model.Trigger{
		Name:      cmd.Name,
		Immediate: cmd.Immediate,
		Frequency: cmd.Frequency,
	})
}

func (t *TriggerAppServImpl) UpdateTrigger(cmd *notification.TriggerUpdateCmd) error {
	return t.gateway.UpdateTrigger(&model.Trigger{
		ID:        cmd.ID,
		Name:      cmd.Name,
		Immediate: cmd.Immediate,
		Frequency: cmd.Frequency,
	})
}
