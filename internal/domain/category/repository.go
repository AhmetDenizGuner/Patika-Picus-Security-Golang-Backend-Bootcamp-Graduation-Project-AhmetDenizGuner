package category

import "gorm.io/gorm"

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) FindAll() (Category, error) {
	var rootCategory Category

	result := r.db.Preload("Category").Where("table_category.id = ?", 1).Find(&rootCategory)

	if result.Error != nil {
		return Category{}, result.Error
	}

	return rootCategory, nil

}

func (r *CategoryRepository) Create(c Category) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
