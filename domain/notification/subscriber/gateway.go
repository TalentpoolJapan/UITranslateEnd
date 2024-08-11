package subscriber

type ExternalJobseekerInfo struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type IGateway interface {
	GetJobseekerInfo(uuid string) (*ExternalJobseekerInfo, error)
}
