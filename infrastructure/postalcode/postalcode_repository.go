package postalcode

import (
	"uitranslate/domain/postalcode"
	"uitranslate/infrastructure"
	"xorm.io/xorm"
)

type MysqlPostalCodeInfoRepository struct {
	DB *xorm.Engine
}

func NewMysqlPostalCodeInfoRepository() *MysqlPostalCodeInfoRepository {
	return &MysqlPostalCodeInfoRepository{
		DB: infrastructure.MysqlDB,
	}
}

type postalCodeInfoPO struct {
	PostalCode      string `xorm:"'postal_code' not null pk varchar(7) comment('邮政编码')"`
	PrefectureKana  string `xorm:"'prefecture_kana' not null varchar(255) comment('都道府县名（片假名）')"`
	CityKana        string `xorm:"'city_kana' not null varchar(255) comment('市区町村名（片假名）')"`
	TownKana        string `xorm:"'town_kana' not null varchar(255) comment('町域名（片假名）')"`
	PrefectureKanji string `xorm:"'prefecture_kanji' not null varchar(255) comment('都道府县名（汉字）')"`
	CityKanji       string `xorm:"'city_kanji' not null varchar(255) comment('市区町村名（汉字）')"`
	TownKanji       string `xorm:"'town_kanji' not null varchar(255) comment('町域名（汉字）')"`
	PrefectureID    int64  `xorm:"'prefecture_id' not null default 0 bigint(11) comment('都道府县id')"`
}

func (p postalCodeInfoPO) convertToEntity() *postalcode.PostalCodeInfo {
	return &postalcode.PostalCodeInfo{
		PostalCode:   p.PostalCode,
		Prefecture:   p.PrefectureKanji,
		City:         p.CityKanji,
		Town:         p.TownKanji,
		PrefectureId: p.PrefectureID,
	}
}

func (r *MysqlPostalCodeInfoRepository) GetPostCodeInfoByPostCode(postCode string) (*postalcode.PostalCodeInfo, error) {
	var postalCodeInfoPO postalCodeInfoPO
	has, err := r.DB.Table("postal_codes_info").Where("postal_code = ?", postCode).Get(&postalCodeInfoPO)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return postalCodeInfoPO.convertToEntity(), nil
}
