package test

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"uitranslate/domain/category"
	"uitranslate/infrastructure"
	"uitranslate/infrastructure/category/repo"
)

func init() {
	infrastructure.InitMysqlDB()
}

func TestCategoryRepository_CreateCategory(t *testing.T) {
	repository := repo.NewCategoryRepository()

	// Create a category
	category := &category.Category{
		// Populate category fields here
	}
	err := repository.CreateCategory(category)
	if err != nil {
		t.Errorf("Failed to create category: %v", err)
	}
	// Additional assertions if needed
}

func TestCategoryRepository_GetCategoryById(t *testing.T) {
	repository := repo.NewCategoryRepository()

	// Assuming there is a category with ID = 1 in the database
	categoryID := int64(1)
	expectedCategoryName := "Some Category"

	category, err := repository.GetCategoryById(categoryID)
	if err != nil {
		t.Errorf("Failed to get category: %v", err)
	}

	if category == nil {
		t.Error("Category not found")
	}

	if category.Name != expectedCategoryName {
		t.Errorf("Expected category name %s, got %s", expectedCategoryName, category.Name)
	}
	// Additional assertions if needed
}
