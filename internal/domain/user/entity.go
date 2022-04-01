package user

import (
	"crypto/sha256"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	PasswordHash string
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
	data := []byte("hello")
	hash := sha256.Sum256(data)
	strHash := fmt.Sprintf("%x", hash)
	u.PasswordHash = strHash
}
