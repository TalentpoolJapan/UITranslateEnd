package topic

type Repository interface {
	GetTopicInfoById(topicId int64) (*BasicInfo, error)
	ListTopicInfo() ([]*BasicInfo, error)
	SaveTopicInfo(topicInfo *BasicInfo) error
	UpdateTopicInfo(topicInfo *BasicInfo) error

	GetTopicTemplateById(templateId int64) (*Template, error)
	ListTopicTemplateByTopicId(topicId int64) ([]*Template, error)
	SaveTopicTemplate(topicTemplate *Template) error
	UpdateTopicTemplate(topicTemplate *Template) error
}
