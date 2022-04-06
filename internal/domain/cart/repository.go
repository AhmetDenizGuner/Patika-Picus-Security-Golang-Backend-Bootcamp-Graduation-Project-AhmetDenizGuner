package cart

import "gorm.io/gorm"

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) FindByID(id int) (Cart, error) {
	var cart Cart

	result := r.db.Preload("Items").Where("id = ?", id).Find(&cart)

	if result.Error != nil {
		return Cart{}, result.Error
	}

	return cart, nil
}

func (r *CartRepository) FindByUserId(userId int) (Cart, error) {
	var cart Cart

	result := r.db.Preload("Items").Where("user_id = ?", userId).Find(&cart)

	if result.Error != nil {
		return Cart{}, result.Error
	}

	return cart, nil
}

func (r *CartRepository) Create(c Cart) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) Update(c Cart) error {
	result := r.db.Save(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
