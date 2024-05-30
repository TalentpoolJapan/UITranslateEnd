package notification

import (
	"errors"
)

type Producer struct {
	EventType EventType
	EventData interface{}
}

func (p *Producer) Publish() error {
	// 1. 查找与事件类型关联的主题
	topic, err := p.findTopicByEventType(p.EventType)
	if err != nil {
		return err
	}

	// 2. 找到配置的触发器
	trigger, err := p.findTriggerByTopic(topic)
	if err != nil {
		return err
	}

	// 3. 根据触发器的配置，将主题交给任务执行器处理
	if trigger.Immediate {
		// 如果触发器设置为立即执行，则立即处理主题
		return p.handleTopic(topic, p.EventData)
	} else {
		// 如果触发器设置为频率执行，则在设置的时间后处理主题

	}

	return nil
}

func (p *Producer) findTopicByEventType(eventType EventType) (Topic, error) {
	// 这里需要实现查找与事件类型关联的主题的逻辑
	// 如果找不到主题，返回错误
	return Topic{}, errors.New("topic not found")
}

func (p *Producer) findTriggerByTopic(topic Topic) (Trigger, error) {
	// 这里需要实现查找与主题关联的触发器的逻辑
	// 如果找不到触发器，返回错误
	return Trigger{}, errors.New("trigger not found")
}

func (p *Producer) handleTopic(topic Topic, eventData interface{}) error {
	// 这里需要实现处理主题的逻辑
	// 这可能包括将主题和事件数据发送给任务执行器等
	return nil
}
