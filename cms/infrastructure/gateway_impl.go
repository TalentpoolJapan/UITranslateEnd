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

func (c CategoryGatewayImpl) QueryCategoryByName(name string) ([]*model.Category, error) {
	wrapper := &repo.QueryWrapper{
		Name: name,
	}
	categories, err := c.repo.DycQuery(wrapper)
	if err != nil {
		return nil, err
	}
	return categories, nil
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
	wrapper := &repo.QueryWrapper{
		ParentId: param.ParentId,
		Name:     param.Name,
	}
	total, categories, err := c.repo.PageCategory(int(param.Page), int(param.PageSize), wrapper)
	if err != nil {
		return 0, nil, err
	}
	return total, categories, nil
}

func NewCategoryGatewayImpl() *CategoryGatewayImpl {
	return &CategoryGatewayImpl{
		repo: repo.NewCategoryRepository(),
	}
}
