package gateway

import (
	"uitranslate/domain/model"
	"uitranslate/domain/param"
)

type ICategoryGateWay interface {
	AddCategory(category *model.Category) error

	UpdateCategory(category *model.Category) error

	DeleteCategory(category *model.Category) error

	PageCategory(param *param.QueryCategoryPage) (int64, []*model.Category, error)

	QueryCategoryByName(name string) ([]*model.Category, error)

	QueryCategoryByParentId(parentId int64) ([]*model.Category, error)
}
