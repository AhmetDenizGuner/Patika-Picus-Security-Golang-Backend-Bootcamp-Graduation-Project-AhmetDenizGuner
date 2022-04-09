package order_item

import "gorm.io/gorm"

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{
		db: db,
	}
}

func (r *OrderItemRepository) DeleteById(id uint) error {
	result := r.db.Delete(&OrderItem{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
