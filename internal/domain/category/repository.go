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

func (r *CategoryRepository) FindByPagination(offset, pageSize int) ([]Category, error) {
	var categories []Category

	result := r.db.Preload("Parent").Offset(offset).Limit(pageSize).Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (r *CategoryRepository) Create(c *Category) error {
	result := r.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CategoryRepository) FindById(id int) (Category, error) {
	var category Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		return Category{}, result.Error
	}
	return category, nil
}

func (r *CategoryRepository) MigrateTable() {
	r.db.AutoMigrate(&Category{})
}
