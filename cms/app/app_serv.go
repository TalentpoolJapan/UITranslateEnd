package app

import "uitranslate/cms/app/dto"

type ICategoryApplicationService interface {
	AddCategory(req dto.AddCategoryReq) error

	UpdateCategory(req dto.UpdateCategoryReq) error

	// query
	PageCategory(req dto.CategoryPageReq) (dto.CategoryPageResp, error)

	CategoryApiData(name string) ([]*dto.CategoryDetailResp, error)
}
