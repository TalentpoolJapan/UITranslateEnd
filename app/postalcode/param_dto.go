package postalcode

type PostalCodeDetailDTO struct {
	PostalCode string `json:"post_code"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
	Town       string `json:"town"`

	PrefectureId int64 `json:"prefecture_id"`
}
