package user

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/api/types"
	"github.com/pkg/errors"
)

type UserService struct {
	repository UserRepository
}

//NewUserService is constructor of UserService
func NewUserService(r UserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}

//SignupService is used for creating new user
func (service *UserService) SignupService(userInfo types.SignupRequest) error {

	//Check email is already used
	_, err := service.repository.FindByEmail(userInfo.Email)

	if err == nil {
		return errors.New("The email is already used for another account!")
	}

	user := NewUser(userInfo.Name, userInfo.Email, userInfo.Password)

	//check validations
	err2 := user.Validate()

	if err2 != nil {
		return err2
	}

	//Add new user to database
	err3 := service.repository.Create(*user)

	if err3 != nil {
		return err3
	}

	return nil
}
