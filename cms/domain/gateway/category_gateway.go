package gateway

import (
	"uitranslate/cms/domain/model"
	"uitranslate/cms/domain/param"
)

type ICategoryGateWay interface {
	AddCategory(category *model.Category) error

	UpdateCategory(category *model.Category) error

	DeleteCategory(category *model.Category) error

	PageCategory(param *param.QueryCategoryPage) (int64, []*model.Category, error)

	QueryCategoryByName(name string) ([]*model.Category, error)
}
