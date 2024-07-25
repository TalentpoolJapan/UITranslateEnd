package gateway

type ExternalUserInfo struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserGateway interface {
	GetUserInfo(uuid string, userType string) (*ExternalUserInfo, error)
}
