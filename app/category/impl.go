package category

import (
	dto2 "uitranslate/app/category/dto"
	category2 "uitranslate/domain/category"
	"uitranslate/infrastructure/category"
)

var (
	Impl ICategoryApplicationService = &CategoryApplicationServiceImpl{}
)

type CategoryApplicationServiceImpl struct {
	gateway category2.ICategoryGateWay
}

func (c *CategoryApplicationServiceImpl) ListCategoryApiDataByName(name string) ([]*dto2.CategoryDetailResp, error) {
	categories, err := c.gateway.ListCategoryByName(name)
	if err != nil {
		return nil, err
	}
	return dto2.ToDtoList(categories), nil
}

func (c *CategoryApplicationServiceImpl) CategoryApiDataById(id int64) (*dto2.CategoryDetailResp, error) {
	category, err := c.gateway.GetCategoryById(id)
	if err != nil {
		return nil, err
	}
	return dto2.ToDto(category), err
}

func (c *CategoryApplicationServiceImpl) AllCategory(parentId int64) ([]*dto2.CategoryDetailResp, error) {
	categories, err := c.gateway.ListCategoryByParentId(parentId)
	if err != nil {
		return nil, err
	}
	return dto2.ToDtoList(categories), nil
}

func (c *CategoryApplicationServiceImpl) PageCategory(req dto2.CategoryPageReq) (dto2.CategoryPageResp, error) {
	totalRow, categories, err := c.gateway.PageCategory(req.ToQuery())
	if err != nil {
		// todo
		return dto2.CategoryPageResp{}, err
	}
	resp := dto2.CategoryPageResp{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    totalRow,
		Data:     dto2.ToDtoList(categories),
	}
	return resp, nil
}

func (c *CategoryApplicationServiceImpl) AddCategory(req dto2.AddCategoryReq) error {
	category := req.ToCategory()
	err := c.gateway.AddCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryApplicationServiceImpl) UpdateCategory(req dto2.UpdateCategoryReq) error {
	category := req.ToCategory()
	err := c.gateway.UpdateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func NewCategoryApplicationService() *CategoryApplicationServiceImpl {
	return &CategoryApplicationServiceImpl{
		gateway: *category.NewCategoryGatewayImpl(),
	}
}
