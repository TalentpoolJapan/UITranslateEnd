package category

type ICategoryGateWay interface {
	AddCategory(category *Category) error

	UpdateCategory(category *Category) error

	DeleteCategory(category *Category) error

	PageCategory(param *QueryCategoryPage) (int64, []*Category, error)

	GetCategoryById(id int64) (*Category, error)

	ListCategoryByName(name string) ([]*Category, error)

	ListCategoryByParentId(parentId int64) ([]*Category, error)
}
