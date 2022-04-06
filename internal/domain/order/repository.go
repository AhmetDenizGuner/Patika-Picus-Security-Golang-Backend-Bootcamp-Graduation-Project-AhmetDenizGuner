package order

import "gorm.io/gorm"

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(c Order) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
