package notification

import (
	"errors"
)

type Task struct {
	Topic    Topic
	Producer Producer
}

func (t *Task) ExecuteTask() error {
	// 1. 获取subscriber信息，检查是否订阅该topic，以及支持的channel（如支持发送email）
	if !t.isSubscribed() {
		return errors.New("subscriber has not subscribed to the topic")
	}

	// 2. 根据topic的模版配置，以及用户支持的channel，获取最合适的模版
	template, err := t.getTemplate()
	if err != nil {
		return err
	}

	// 3. 内容构建
	// 3.1 根据模版获取 producer & subscriber的模版参数
	params, err := t.getTemplateParams()
	if err != nil {
		return err
	}

	// 3.2 加载、渲染模版
	content, err := t.renderTemplate(template, params)
	if err != nil {
		return err
	}

	// 4. 根据channel发送消息到subscriber
	err = t.sendMessage(content)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) isSubscribed() bool {
	// 这里需要实现检查subscriber是否订阅了topic的逻辑
	return false
}

func (t *Task) getTemplate() (string, error) {
	// 这里需要实现根据topic的模版配置，以及用户支持的channel，获取最合适的模版的逻辑
	return "", errors.New("template not found")
}

func (t *Task) getTemplateParams() (map[string]interface{}, error) {
	// 这里需要实现根据模版获取 producer & subscriber的模版参数的逻辑
	return nil, errors.New("failed to get template params")
}

func (t *Task) renderTemplate(template string, params map[string]interface{}) (string, error) {
	// 这里需要实现加载和渲染模版的逻辑
	return "", errors.New("failed to render template")
}

func (t *Task) sendMessage(content string) error {
	// 这里需要实现根据channel发送消息到subscriber的逻辑
	return errors.New("failed to send message")
}
