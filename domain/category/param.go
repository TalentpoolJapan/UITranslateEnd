package category

type QueryCategoryPage struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}
