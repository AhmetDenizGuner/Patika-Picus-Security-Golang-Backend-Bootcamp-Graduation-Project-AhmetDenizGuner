package user

import (
	"crypto/sha256"
	"fmt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/role"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `validate:"min=3"`
	Email        string `validate:"email"`
	Password     string `validate:"min=3"`
	PasswordHash string
	Roles        []role.Role `gorm:"many2many:user_role;"`
}

func NewUser(name, email, password string) *User {
	user := &User{
		Name:         name,
		Email:        email,
		Password:     password,
		PasswordHash: "",
	}

	user.SetPasswordHash()

	return user
}

func (u *User) SetPasswordHash() {
	data := []byte(u.Password)
	hash := sha256.Sum256(data)
	strHash := fmt.Sprintf("%x", hash)
	u.PasswordHash = strHash
}

func (u *User) Validate() error {

	validate := validator.New()
	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil

}
