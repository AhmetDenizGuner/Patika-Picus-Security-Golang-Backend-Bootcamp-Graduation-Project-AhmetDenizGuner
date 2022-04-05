package product

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/category"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	StockCode     string
	StockQuantity int
	Price         float64
	Description   string
	CategoryID    uint
	Category      category.Category `gorm:"foreignKey:CategoryID"`
}

func NewProduct(name, stockCode string, stockQuantity int, price float64, description string, categoryId uint) *Product {
	product := &Product{
		Name:          name,
		StockCode:     stockCode,
		StockQuantity: stockQuantity,
		Price:         price,
		Description:   description,
		CategoryID:    categoryId,
	}

	return product
}

func (p *Product) Update(name, stockCode string, stockQuantity int, price float64, description string, categoryId uint) {
	p.Name = name
	p.StockCode = stockCode
	p.StockQuantity = stockQuantity
	p.Price = price
	p.Description = description
	p.CategoryID = categoryId
}
