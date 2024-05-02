package app

import (
	"uitranslate/app/dto"
	"uitranslate/domain/gateway"
	"uitranslate/infrastructure"
)

var (
	Impl ICategoryApplicationService = &CategoryApplicationServiceImpl{}
)

type CategoryApplicationServiceImpl struct {
	gateway gateway.ICategoryGateWay
}

func (c *CategoryApplicationServiceImpl) ListCategoryApiDataByName(name string) ([]*dto.CategoryDetailResp, error) {
	categories, err := c.gateway.QueryCategoryByName(name)
	if err != nil {
		return nil, err
	}
	return dto.ToDtoList(categories), nil
}

func (c *CategoryApplicationServiceImpl) CategoryApiDataById(id int64) (*dto.CategoryDetailResp, error) {
	category, err := c.gateway.QueryCategoryById(id)
	if err != nil {
		return nil, err
	}
	return dto.ToDto(category), err
}

func (c *CategoryApplicationServiceImpl) AllCategory(parentId int64) ([]*dto.CategoryDetailResp, error) {
	categories, err := c.gateway.QueryCategoryByParentId(parentId)
	if err != nil {
		return nil, err
	}
	return dto.ToDtoList(categories), nil
}

func (c *CategoryApplicationServiceImpl) PageCategory(req dto.CategoryPageReq) (dto.CategoryPageResp, error) {
	totalRow, categories, err := c.gateway.PageCategory(req.ToQuery())
	if err != nil {
		// todo
		return dto.CategoryPageResp{}, err
	}
	resp := dto.CategoryPageResp{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    totalRow,
		Data:     dto.ToDtoList(categories),
	}
	return resp, nil
}

func (c *CategoryApplicationServiceImpl) AddCategory(req dto.AddCategoryReq) error {
	category := req.ToCategory()
	err := c.gateway.AddCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryApplicationServiceImpl) UpdateCategory(req dto.UpdateCategoryReq) error {
	category := req.ToCategory()
	err := c.gateway.UpdateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func NewCategoryApplicationService() *CategoryApplicationServiceImpl {
	return &CategoryApplicationServiceImpl{
		gateway: *infrastructure.NewCategoryGatewayImpl(),
	}
}
