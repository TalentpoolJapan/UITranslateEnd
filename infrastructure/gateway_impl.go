package infrastructure

import (
	"uitranslate/domain/gateway"
	"uitranslate/domain/model"
	"uitranslate/domain/param"
	"uitranslate/infrastructure/repo"
)

var (
	Impl gateway.ICategoryGateWay = &CategoryGatewayImpl{}
)

type CategoryGatewayImpl struct {
	repo *repo.CategoryRepository
}

func (c CategoryGatewayImpl) QueryCategoryById(id int64) (*model.Category, error) {
	category, err := c.repo.GetCategoryById(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategoryGatewayImpl) QueryCategoryByParentId(parentId int64) ([]*model.Category, error) {
	wrapper := &repo.QueryWrapper{
		ParentId: parentId,
	}
	categories, err := c.repo.DycQuery(wrapper)
	if err != nil {
		return nil, err
	}
	return categories, nil
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
	// 确认id，每个层级最多1000个
	var newCategoryId int64
	wrapper := &repo.QueryWrapper{
		ParentId: category.ParentId,
	}
	categoryGroup, _ := c.repo.DycQuery(wrapper)
	if len(categoryGroup) == 0 {
		newCategoryId = category.ParentId * 1000
	} else {
		// 获取category group中id最大的category
		// 初始化 maxCategory 为数组的第一个元素
		maxCategory := categoryGroup[0]
		// 遍历数组，比较每个 Category 的 ID，更新最大的 Category
		for _, category := range categoryGroup {
			if category.ID > maxCategory.ID {
				maxCategory = category
			}
		}
		newCategoryId = maxCategory.ID + 1
	}

	category.ID = newCategoryId
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
