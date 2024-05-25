package category

import (
	dto2 "uitranslate/app/category/dto"
	category2 "uitranslate/domain/category"
	"uitranslate/infrastructure/category"
)

var (
	Impl ICategoryApplicationService = &CategoryApplicationServiceImpl{
		gateway: category.Impl,
	}
)

type CategoryApplicationServiceImpl struct {
	gateway category2.Gateway
}

func (serv *CategoryApplicationServiceImpl) ListCategoryByParentName(name string) ([]*dto2.CategoryDetailResp, error) {
	allCategories, err := serv.gateway.ListCategoryByParentName(name)
	if err != nil {
		return nil, err
	}
	// filter status
	var publishCategories []*category2.Category
	for _, c := range allCategories {
		if c.Status == category2.Published {
			publishCategories = append(publishCategories, c)
		}
	}
	return dto2.ToDtoList(publishCategories), nil
}

func (serv *CategoryApplicationServiceImpl) CategoryApiDataById(id int64) (*dto2.CategoryDetailResp, error) {
	category, err := serv.gateway.GetCategoryById(id)
	if err != nil {
		return nil, err
	}
	return dto2.ToDto(category), err
}

func (serv *CategoryApplicationServiceImpl) AllCategory(parentId int64) ([]*dto2.CategoryDetailResp, error) {
	categories, err := serv.gateway.ListCategoryByParentId(parentId)
	if err != nil {
		return nil, err
	}
	return dto2.ToDtoList(categories), nil
}

func (serv *CategoryApplicationServiceImpl) PageCategory(req dto2.CategoryPageReq) (dto2.CategoryPageResp, error) {
	totalRow, categories, err := serv.gateway.PageCategory(req.ToQuery())
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

func (serv *CategoryApplicationServiceImpl) AddCategory(req dto2.AddCategoryReq) error {
	category := req.ToCategory()
	err := serv.gateway.AddCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (serv *CategoryApplicationServiceImpl) UpdateCategory(req dto2.UpdateCategoryReq) error {
	category := req.ToCategory()
	err := serv.gateway.UpdateCategory(category)
	if err != nil {
		return err
	}
	return nil
}
