package repo

import (
	"uitranslate/cms/domain/model"
	"uitranslate/cms/infrastructure/repo/po"
)

var categoryTableName = "talentpool_category"

type CategoryRepository struct {
}

type QueryWrapper struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (repo *CategoryRepository) CreateCategory(category *model.Category) (err error) {
	_, err = MysqlDB.Table(categoryTableName).Insert(po.ToPO(category))
	return err
}

func (repo *CategoryRepository) UpdateCategory(category *model.Category) (err error) {
	_, err = MysqlDB.Table(categoryTableName).Where("id = ?", category.ID).Update(po.ToPO(category))
	return err
}

func (repo *CategoryRepository) DeleteCategory(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (repo *CategoryRepository) PageCategory(page int, pageSize int, wrapper *QueryWrapper) (int64, []*model.Category, error) {
	var categories []*po.CategoryPO
	//total, err := MysqlDB.Table(categoryTableName).Limit(pageSize, (page-1)*pageSize).FindAndCount(&categories)
	table := MysqlDB.Table(categoryTableName)
	if wrapper.ParentId != 0 {
		table.Where("parent_id = ?", wrapper.ParentId)
	}
	if wrapper.Name != "" {
		table.Where("name = ?", wrapper.Name)
	}
	total, err := table.Limit(pageSize, (page-1)*pageSize).FindAndCount(&categories)
	if err != nil {
		return 0, nil, err
	}

	return total, po.ToEntityList(categories), nil
}
