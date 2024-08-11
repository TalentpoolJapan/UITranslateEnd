package subscriber

type Type string

const (
	JobSeeker  = Type("job_seeker")
	Enterprise = Type("enterprise")
)

// aggreate
type Subscriber struct {
	Uuid           string
	Type           Type
	Name           string
	Email          string
	AcceptChannels []string

	TopicIds []int64
}

func (s *Subscriber) alreadySubscribe(topicId int64) bool {
	for _, id := range s.TopicIds {
		if id == topicId {
			return true
		}
	}
	return false
}
