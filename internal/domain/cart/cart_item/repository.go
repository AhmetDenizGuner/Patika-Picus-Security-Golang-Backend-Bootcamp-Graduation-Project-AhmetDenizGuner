package cart_item

import "gorm.io/gorm"

type CartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) *CartItemRepository {
	return &CartItemRepository{
		db: db,
	}
}

func (r *CartItemRepository) Update(c *CartItem) error {
	result := r.db.Save(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartItemRepository) DeleteById(id uint) error {
	result := r.db.Delete(&CartItem{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
