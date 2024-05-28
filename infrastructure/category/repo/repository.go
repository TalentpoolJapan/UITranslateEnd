package repo

import (
	"uitranslate/domain/category"
	"uitranslate/infrastructure"
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

func (repo *CategoryRepository) CreateCategory(category *category.Category) (err error) {
	_, err = infrastructure.MysqlDB.Table(categoryTableName).Insert(ToPO(category))
	return err
}

func (repo *CategoryRepository) UpdateCategory(category *category.Category) (err error) {
	_, err = infrastructure.MysqlDB.Table(categoryTableName).Where("id = ?", category.ID).Update(ToPO(category))
	return err
}

func (repo *CategoryRepository) DeleteCategory(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (repo *CategoryRepository) GetCategoryById(id int64) (*category.Category, error) {
	var category []*CategoryPO
	err := infrastructure.MysqlDB.Table(categoryTableName).Where("id = ?", id).Find(&category)
	if err != nil || len(category) == 0 {
		return nil, err
	}
	return category[0].ToEntity(), nil
}

func (repo *CategoryRepository) PageCategory(page int, pageSize int, wrapper *QueryWrapper) (int64, []*category.Category, error) {
	var categories []*CategoryPO
	//total, err := MysqlDB.Table(categoryTableName).Limit(pageSize, (page-1)*pageSize).FindAndCount(&categories)
	table := infrastructure.MysqlDB.Table(categoryTableName)
	table.Where("parent_id = ?", wrapper.ParentId)
	if wrapper.Name != "" {
		table.Where("name = ?", wrapper.Name)
	}
	total, err := table.Desc("sort_order").Limit(pageSize, (page-1)*pageSize).FindAndCount(&categories)
	if err != nil {
		return 0, nil, err
	}

	return total, ToEntityList(categories), nil
}

func (repo *CategoryRepository) DycQuery(wrapper *QueryWrapper) ([]*category.Category, error) {
	var categories []*CategoryPO
	//total, err := MysqlDB.Table(categoryTableName).Limit(pageSize, (page-1)*pageSize).FindAndCount(&categories)
	table := infrastructure.MysqlDB.Table(categoryTableName)
	table.Where("parent_id = ?", wrapper.ParentId)
	if wrapper.Name != "" {
		table.Where("name = ?", wrapper.Name)
	}
	err := table.Desc("sort_order").Find(&categories)
	if err != nil {
		return nil, err
	}

	return ToEntityList(categories), nil
}
