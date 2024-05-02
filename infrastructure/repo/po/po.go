package po

import (
	"time"
	"uitranslate/domain/model"
)

type CategoryPO struct {
	Id         int64     `json:"id" db:"id" xorm:"pk"`
	Name       string    `json:"name"`
	ParentId   int64     `json:"parent_id"`
	Tag        string    `json:"tag"`
	SortOrder  int64     `json:"sort_order"`
	NameEn     string    `json:"name_en"`
	NameJa     string    `json:"name_ja"`
	Status     int       `json:"status"`
	UpdateTime time.Time `json:"update_time"`
}

func ToPO(entity *model.Category) *CategoryPO {
	return &CategoryPO{
		Id:         entity.ID,
		Name:       entity.Name,
		ParentId:   entity.ParentId,
		Tag:        entity.Tag,
		SortOrder:  entity.SortOrder,
		NameEn:     entity.NameEn,
		NameJa:     entity.NameJa,
		Status:     int(entity.Status),
		UpdateTime: entity.UpdateTime,
	}
}

func (po *CategoryPO) ToEntity() *model.Category {
	return &model.Category{
		ID:         po.Id,
		Name:       po.Name,
		ParentId:   po.ParentId,
		Tag:        po.Tag,
		SortOrder:  po.SortOrder,
		NameEn:     po.NameEn,
		NameJa:     po.NameJa,
		Status:     model.Status(po.Status),
		UpdateTime: po.UpdateTime,
	}
}

func ToEntityList(poList []*CategoryPO) []*model.Category {
	var list []*model.Category
	for _, po := range poList {
		list = append(list, po.ToEntity())
	}
	return list
}
