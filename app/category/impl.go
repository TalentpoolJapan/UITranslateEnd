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

// api用的接口
func (serv *CategoryApplicationServiceImpl) ListCategoryByParentName(name string, language string) ([]*dto2.CategoryDetailResp, error) {
	allCategories, err := serv.gateway.ListCategoryByParentName(name)
	if err != nil {
		return nil, err
	}
	// filter status
	var publishCategories []*dto2.CategoryDetailResp
	// fixme 缓存
	for _, c := range allCategories {
		if c.Status == category2.Published {
			publishCategories = append(publishCategories, serv.buildCategoryDetailResp(c, language))
		}
	}

	return publishCategories, nil
}

func (serv *CategoryApplicationServiceImpl) buildCategoryDetailResp(category *category2.Category, language string) *dto2.CategoryDetailResp {
	detailResp := dto2.ToDto(category)
	if language == "japanese" {
		detailResp.Name = category.NameJa
	} else {
		detailResp.Name = category.NameEn
	}
	children, err := serv.gateway.ListCategoryByParentId(category.ID)
	if err == nil && len(children) > 1 {
		for _, c := range children {
			if c.Status == category2.Published {
				detailResp.Children = append(detailResp.Children, serv.buildCategoryDetailResp(c, language))
			}
		}
	}
	return detailResp
}

func (serv *CategoryApplicationServiceImpl) CategoryApiDataById(id int64) (*dto2.CategoryDetailResp, error) {
	category, err := serv.gateway.GetCategoryById(id)
	if err != nil {
		return nil, err
	}
	return dto2.ToDto(category), err
}

func (serv *CategoryApplicationServiceImpl) AllCategoryByParentId(parentId int64) ([]*dto2.CategoryDetailResp, error) {
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
