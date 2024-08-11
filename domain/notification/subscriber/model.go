package subscriber

type Type string

const (
	JobSeeker  = Type("job_seeker")
	Enterprise = Type("enterprise")
)

// aggreate
type Subscriber struct {
	Uuid     string
	Type     Type
	TopicIds []int64

	Name           string
	Email          string
	AcceptChannels []string
}

func (s *Subscriber) alreadySubscribe(topicId int64) bool {
	for _, id := range s.TopicIds {
		if id == topicId {
			return true
		}
	}
	return false
}

func (s *Subscriber) subscribe(topicId int64) bool {
	if s.alreadySubscribe(topicId) {
		return false
	}
	s.TopicIds = append(s.TopicIds, topicId)
	return true
}

func (s *Subscriber) unsubscribe(topicId int64) bool {
	for i, id := range s.TopicIds {
		if id == topicId {
			s.TopicIds = append(s.TopicIds[:i], s.TopicIds[i+1:]...)
			return true
		}
	}
	return false
}
