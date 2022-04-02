package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	var user User
	result := r.db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r *UserRepository) Create(u User) error {
	result := r.db.Create(u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

