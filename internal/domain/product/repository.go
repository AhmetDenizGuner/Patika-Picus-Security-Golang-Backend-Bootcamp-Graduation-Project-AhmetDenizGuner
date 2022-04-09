package product

import (
	"gorm.io/gorm"
	"strings"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) FindByPagination(offset, pageSize int) ([]Product, error) {
	var products []Product

	result := r.db.Preload("Category").Offset(offset).Limit(pageSize).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *ProductRepository) FindByPaginationAndKey(offset, pageSize int, key string) ([]Product, error) {
	var products []Product
	key = "%" + key + "%"
	result := r.db.Preload("Category").Where("upper(table_product.Name) LIKE ? OR upper(table_product.stock_code) LIKE ? OR upper(table_product.description) LIKE ?", strings.ToUpper(key), strings.ToUpper(key), strings.ToUpper(key)).Offset(offset).Limit(pageSize).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *ProductRepository) FindByStockCode(code string) (Product, error) {
	var product Product

	result := r.db.Where("stock_code = ?", code).First(&product)

	if result.Error != nil {
		return Product{}, result.Error
	}

	return product, nil
}

func (r *ProductRepository) Create(p *Product) error {
	result := r.db.Create(&p)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepository) FindByID(id int) (Product, error) {
	var product Product

	result := r.db.Where("id = ?", id).First(&product)

	if result.Error != nil {
		return Product{}, result.Error
	}

	return product, nil
}

func (r *ProductRepository) DeleteById(id uint) error {
	result := r.db.Delete(&Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepository) Update(p Product) error {
	result := r.db.Save(p)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepository) MigrateTable() {
	r.db.AutoMigrate(&Product{})
}
