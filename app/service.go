package app

import (
	"uitranslate/app/dto"
)

type ICategoryApplicationService interface {
	AddCategory(req dto.AddCategoryReq) error

	UpdateCategory(req dto.UpdateCategoryReq) error

	// query
	AllCategory(parentId int64) ([]*dto.CategoryDetailResp, error)

	PageCategory(req dto.CategoryPageReq) (dto.CategoryPageResp, error)

	ListCategoryApiDataByName(name string) ([]*dto.CategoryDetailResp, error)

	CategoryApiDataById(id int64) (*dto.CategoryDetailResp, error)
}
