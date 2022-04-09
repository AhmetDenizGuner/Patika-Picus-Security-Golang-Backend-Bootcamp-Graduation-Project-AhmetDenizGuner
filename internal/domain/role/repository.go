package role

import "gorm.io/gorm"

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) MigrateTable() {
	r.db.AutoMigrate(&Role{})
}

func (r *RoleRepository) Create(role *Role) error {

	result := r.db.Create(&role)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *RoleRepository) FindByName(name string) (Role, error) {
	var product Role

	result := r.db.Where("id = ?", name).Find(&product)

	if result.Error != nil {
		return Role{}, result.Error
	}

	return product, nil
}
