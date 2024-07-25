package category

import (
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
	"uitranslate/domain/category"
	"uitranslate/infrastructure/category/repo"
)

var (
	Impl category.Gateway = &GatewayImpl{
		repo:  repo.NewCategoryRepository(),
		cache: cache.New(1*time.Hour, 1*time.Hour),
	}
)

type GatewayImpl struct {
	repo  *repo.CategoryRepository
	cache *cache.Cache
}

func (c GatewayImpl) GetCategoryById(id int64) (*category.Category, error) {
	categoryById, err := c.repo.GetCategoryById(id)
	if err != nil {
		return nil, err
	}
	return categoryById, nil
}

func (c GatewayImpl) ListCategoryByParentId(parentId int64) ([]*category.Category, error) {
	cacheKey := "categories_" + strconv.FormatInt(parentId, 10)
	if categories, found := c.cache.Get(cacheKey); found {
		return categories.([]*category.Category), nil
	}

	wrapper := &repo.QueryWrapper{
		ParentId: parentId,
	}
	categories, err := c.repo.DycQuery(wrapper)
	if err != nil {
		return nil, err
	}
	c.cache.Set(cacheKey, categories, 1*time.Hour)
	return categories, nil
}

func (c GatewayImpl) ListCategoryByParentName(parentName string) ([]*category.Category, error) {
	parentCategories, err := c.repo.DycQuery(&repo.QueryWrapper{
		Name: parentName,
	})
	if err != nil {
		return nil, err
	}
	if len(parentCategories) == 0 {
		return nil, nil
	}
	parentId := parentCategories[0].ID
	categories, err := c.repo.DycQuery(&repo.QueryWrapper{
		ParentId: parentId,
	})
	return categories, nil
}

func (c GatewayImpl) AddCategory(category *category.Category) error {
	// 确认id，每个层级最多1000个
	var newCategoryId int64
	wrapper := &repo.QueryWrapper{
		ParentId: category.ParentId,
	}
	categoryGroup, _ := c.repo.DycQuery(wrapper)
	if len(categoryGroup) == 0 {
		newCategoryId = category.ParentId*1000 + 1
	} else {
		// 获取category group中id最大的category
		// 初始化 maxCategory 为数组的第一个元素
		maxCategory := categoryGroup[0]
		// 遍历数组，比较每个 Category 的 Id，更新最大的 Category
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

func (c GatewayImpl) UpdateCategory(category *category.Category) error {
	err := c.repo.UpdateCategory(category)
	return err
}

func (c GatewayImpl) DeleteCategory(category *category.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c GatewayImpl) PageCategory(param *category.QueryCategoryPage) (int64, []*category.Category, error) {
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
