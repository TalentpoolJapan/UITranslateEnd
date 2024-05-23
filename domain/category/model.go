package category

import (
	"time"
)

type Status int

const (
	Published Status = iota
	Unpublished
)

type Category struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	ParentId   int64     `json:"parent_id"`
	Tag        string    `json:"tag"` // 兼容原来的 tag
	SortOrder  int64     `json:"sort_order"`
	NameEn     string    `json:"name_en"`
	NameJa     string    `json:"name_ja"`
	Status     Status    `json:"status"`
	UpdateTime time.Time `json:"update_time"`
}
