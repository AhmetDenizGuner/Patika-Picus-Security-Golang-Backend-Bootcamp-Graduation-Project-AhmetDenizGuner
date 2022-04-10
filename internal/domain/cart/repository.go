package cart

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) FindByUserId(userId int) (Cart, error) {
	var cart Cart

	result := r.db.Preload("Items").Preload("Items.Product").Where("user_id = ?", userId).First(&cart)

	if result.Error != nil {
		return Cart{}, result.Error
	}

	return cart, nil
}

func (r *CartRepository) Create(c *Cart) error {
	result := r.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) Update(c *Cart) error {
	result := r.db.Save(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) MigrateTable() {
	r.db.AutoMigrate(&Cart{})
	r.db.AutoMigrate(&cart_item.CartItem{})
}
