package category

import (
	dto2 "uitranslate/app/category/dto"
)

type ICategoryApplicationService interface {
	AddCategory(req dto2.AddCategoryReq) error

	UpdateCategory(req dto2.UpdateCategoryReq) error

	// query
	AllCategoryByParentId(parentId int64) ([]*dto2.CategoryDetailResp, error)

	PageCategory(req dto2.CategoryPageReq) (dto2.CategoryPageResp, error)

	ListCategoryByParentName(name string) ([]*dto2.CategoryDetailResp, error)

	CategoryApiDataById(id int64) (*dto2.CategoryDetailResp, error)
}
