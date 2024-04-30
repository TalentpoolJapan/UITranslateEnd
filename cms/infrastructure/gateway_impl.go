package infrastructure

import (
	"uitranslate/cms/domain/gateway"
	"uitranslate/cms/domain/model"
	"uitranslate/cms/domain/param"
	"uitranslate/cms/infrastructure/repo"
)

var (
	Impl gateway.ICategoryGateWay = &CategoryGatewayImpl{}
)

type CategoryGatewayImpl struct {
	repo *repo.CategoryRepository
}

func NewCategoryGatewayImpl() *CategoryGatewayImpl {
	return &CategoryGatewayImpl{
		repo: repo.NewCategoryRepository(),
	}
}

func (c CategoryGatewayImpl) AddCategory(category *model.Category) error {
	err := c.repo.CreateCategory(category)
	return err
}

func (c CategoryGatewayImpl) UpdateCategory(category *model.Category) error {
	err := c.repo.UpdateCategory(category)
	return err
}

func (c CategoryGatewayImpl) DeleteCategory(category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c CategoryGatewayImpl) PageCategory(param *param.QueryCategoryPage) (int64, []*model.Category, error) {
	total, categories, err := c.repo.PageCategory(int(param.Page), int(param.PageSize))
	if err != nil {
		return 0, nil, err
	}
	return total, categories, nil
}
