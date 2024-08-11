package model

import (
	"errors"
	"uitranslate/domain/notification"
	"uitranslate/domain/notification/subscriber"
	"uitranslate/domain/notification/topic"
)

type Task struct {
	TopicId  int64
	Producer Producer
}

func (t *Task) ExecuteTask() error {
	topic, err := t.getTopic()
	if err != nil {
		return err
	}

	subscribers := querySubscriber(topic.TopicInfo)
	if len(subscribers) == 0 {
		return errors.New("no subscriber found")
	}

	for _, subscriber := range subscribers {
		err := t.sendToSubscriber(subscriber, topic)
		if err != nil {
			return err
		}
	}

	return nil
}

func querySubscriber(topic topic.TopicInfo) []*subscriber.Subscriber {
	// todo 这里需要实现查询订阅了topic的subscriber的逻辑
	return nil
}

func (t *Task) getTemplateParams() (map[string]interface{}, error) {
	// 这里需要实现根据模版获取 producer & subscriber的模版参数的逻辑
	return nil, errors.New("failed to get template params")
}

func renderTemplate() (string, error) {
	// todo 这里需要实现加载和渲染模版的逻辑
	return "", errors.New("failed to render template")
}

func (t *Task) getTopic() (topic.AggregateTopic, error) {
	// todo 这里需要实现获取topic信息的逻辑
	return topic.AggregateTopic{}, nil
}

func (t *Task) sendToSubscriber(subscriber *subscriber.Subscriber, topic topic.AggregateTopic) error {
	// todo 这里需要实现将消息发送给subscriber的逻辑
	channels := subscriber.AcceptChannels
	if len(channels) == 0 {
		return errors.New("no channel found")
	}

	for _, channel := range channels {
		template, err := topic.SelectTemplate(notification.Channel(channel))
		if err != nil {
			// todo log
			continue
		}
		doSent(template, subscriber)
	}
	return nil
}

func doSent(template *topic.TopicTemplate, subscriber *subscriber.Subscriber) error {
	// todo
	// 3.2 加载、渲染模版
	//content, err := renderTemplate()
	//if err != nil {
	//	return err
	//}

	// 4. 根据channel发送消息到subscriber
	return nil
}
