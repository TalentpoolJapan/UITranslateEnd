package dto

import (
	"uitranslate/domain/category"
)

type CategoryDetailResp struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ParentId  int64  `json:"parent_id"`
	Tag       string `json:"tag"` // 兼容原来的 tag
	SortOrder int64  `json:"sort_order"`
	NameEn    string `json:"name_en"`
	NameJa    string `json:"name_ja"`
	Status    int    `json:"status"`
}

type CategoryPageResp struct {
	Total    int64                 `json:"total"`
	PageSize int64                 `json:"page_size"`
	Page     int64                 `json:"page"`
	Data     []*CategoryDetailResp `json:"data"`
}

func ToDto(category *category.Category) *CategoryDetailResp {
	dto := &CategoryDetailResp{
		ID:        category.ID,
		Name:      category.Name,
		ParentId:  category.ParentId,
		Tag:       category.Tag,
		SortOrder: category.SortOrder,
		NameEn:    category.NameEn,
		NameJa:    category.NameJa,
		Status:    int(category.Status),
	}
	return dto
}

func ToDtoList(categories []*category.Category) []*CategoryDetailResp {
	var arr []*CategoryDetailResp

	for _, category := range categories {
		arr = append(arr, ToDto(category))
	}
	return arr
}
