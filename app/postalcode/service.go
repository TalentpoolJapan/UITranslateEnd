package postalcode

import (
	application "uitranslate/app"
	"uitranslate/domain/postalcode"
)

type PostalCodeApplicationService interface {
	GetPostalCodeDetail(qry *PostalCodeDetailQry) application.SingleResp[PostalCodeDetailDTO]
}

type postalCodeApplicationService struct {
	postalCodeInfoRepository postalcode.PostalCodeInfoRepository
}

func NewPostalCodeApplicationService(postalCodeInfoRepository postalcode.PostalCodeInfoRepository) PostalCodeApplicationService {
	return &postalCodeApplicationService{
		postalCodeInfoRepository: postalCodeInfoRepository,
	}
}

func (p postalCodeApplicationService) GetPostalCodeDetail(qry *PostalCodeDetailQry) application.SingleResp[PostalCodeDetailDTO] {
	postCodeInfo, err := p.postalCodeInfoRepository.GetPostCodeInfoByPostCode(qry.PostalCode)
	if err != nil {
		return application.SingleRespFail[PostalCodeDetailDTO](err.Error())
	}

	if postCodeInfo == nil {
		return application.SingleRespOk[PostalCodeDetailDTO]()
	}

	return application.SingleRespOf[PostalCodeDetailDTO](PostalCodeDetailDTO{
		PostalCode:   postCodeInfo.PostalCode,
		Prefecture:   postCodeInfo.Prefecture,
		City:         postCodeInfo.City,
		Town:         postCodeInfo.Town,
		PrefectureId: postCodeInfo.PrefectureId,
	}, "success")
}
