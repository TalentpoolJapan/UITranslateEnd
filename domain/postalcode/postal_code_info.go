package postalcode

type PostalCodeInfo struct {
	PostalCode string
	Prefecture string
	City       string
	Town       string

	PrefectureId int64
}

type PostalCodeInfoRepository interface {
	GetPostCodeInfoByPostCode(postCode string) (*PostalCodeInfo, error)
}
