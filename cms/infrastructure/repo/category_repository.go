package repo

import (
	"uitranslate/cms/domain/model"
	"uitranslate/cms/infrastructure/repo/po"
)

var categoryTableName = "talentpool_category"

type CategoryRepository struct {
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (repo *CategoryRepository) CreateCategory(category *model.Category) (err error) {
	_, err = MysqlDB.Table(categoryTableName).Insert(&category)
	return err
}

func (repo *CategoryRepository) UpdateCategory(category *model.Category) (err error) {
	_, err = MysqlDB.Table(categoryTableName).Update(&category)
	return err
}

func (repo *CategoryRepository) DeleteCategory(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (repo *CategoryRepository) PageCategory(page int, pageSize int) (int64, []*model.Category, error) {
	var categories []*po.CategoryPO
	total, err := MysqlDB.Table(categoryTableName).Limit(pageSize, (page-1)*pageSize).FindAndCount(&categories)
	if err != nil {
		return 0, nil, err
	}

	return total, po.ToEntityList(categories), nil
}
