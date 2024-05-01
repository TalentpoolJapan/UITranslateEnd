package dto

import (
	"time"
	"uitranslate/cms/domain/model"
	"uitranslate/cms/domain/param"
)

type AddCategoryReq struct {
	Name      string `json:"name" binding:"required"`
	ParentId  int64  `json:"parent_id"`
	Tag       string `json:"tag"`
	SortOrder int64  `json:"sort_order"`
	NameEn    string `json:"name_en" binding:"required"`
	NameJa    string `json:"name_ja" binding:"required"`
}

type UpdateCategoryReq struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ParentId  int64  `json:"parent_id"`
	Tag       string `json:"tag"` // 兼容原来的 tag
	SortOrder int64  `json:"sort_order"`
	NameEn    string `json:"name_en"`
	NameJa    string `json:"name_ja"`
	Status    int    `json:"status"`
}

type CategoryPageReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

func (dto *AddCategoryReq) ToCategory() *model.Category {
	category := &model.Category{
		Name:       dto.Name,
		ParentId:   dto.ParentId,
		Tag:        dto.Tag,
		SortOrder:  dto.SortOrder,
		NameEn:     dto.NameEn,
		NameJa:     dto.NameJa,
		UpdateTime: time.Now(),
	}
	return category
}

func (dto *UpdateCategoryReq) ToCategory() *model.Category {
	category := &model.Category{
		ID:        dto.ID,
		Name:      dto.Name,
		ParentId:  dto.ParentId,
		Tag:       dto.Tag,
		SortOrder: dto.SortOrder,
		NameEn:    dto.NameEn,
		NameJa:    dto.NameJa,
		Status:    model.Status(dto.Status),
	}
	return category
}

func (dto *CategoryPageReq) ToQuery() *param.QueryCategoryPage {
	query := &param.QueryCategoryPage{
		Page:     dto.Page,
		PageSize: dto.PageSize,
		ParentId: dto.ParentId,
		Name:     dto.Name,
	}
	return query
}
