package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"uitranslate/app/postalcode"
	postalcode2 "uitranslate/infrastructure/postalcode"
)

var (
	postalCodeAppService = postalcode.NewPostalCodeApplicationService(postalcode2.NewMysqlPostalCodeInfoRepository())
)

func RegisterPostalCodeHandler(engine *gin.Engine) {
	engine.GET("/api/postal_code", getPostalCodeInfoByCode)
}

func getPostalCodeInfoByCode(context *gin.Context) {
	var qry postalcode.PostalCodeDetailQry
	language := context.GetHeader("TalentPool-Language")
	if language == "" {
		language = "japanese"
	}
	postalCode := context.Query("postal_code")
	if postalCode == "" {
		context.JSON(http.StatusBadRequest, NewApiRestResult(RestResult{Code: -1, Message: "postal code empty"}))
		return
	}

	qry.PostalCode = postalCode
	qry.Language = language

	postalCodeInfoResp := postalCodeAppService.GetPostalCodeDetail(&qry)
	context.JSON(http.StatusOK, NewApiRestResult(RestResult{Code: 0, Message: postalCodeInfoResp.Msg, Data: postalCodeInfoResp.Data}))
}
