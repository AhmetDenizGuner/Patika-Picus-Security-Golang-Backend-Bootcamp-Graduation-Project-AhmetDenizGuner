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

func (r *OrderRepository) FindAllByUser(userID int) ([]Order, error) {

	var orders []Order

	result := r.db.Preload("Items").Preload("Items.Product").Where("user_id = ? ", userID).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil

}

func (r *OrderRepository) FindByID(id int) (Order, error) {
	var order Order

	result := r.db.Preload("Items").Preload("Items.Product").Where("id = ?", id).First(&order)

	if result.Error != nil {
		return Order{}, result.Error
	}

	return order, nil
}

func (r *OrderRepository) DeleteById(id uint) error {
	result := r.db.Delete(&Order{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
