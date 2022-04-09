package user

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/api/types"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/role"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/redis"
	"log"
)

type UserService struct {
	repository     UserRepository
	roleRepository role.RoleRepository
}

//NewUserService is constructor of UserService
func NewUserService(r UserRepository, roleRepository role.RoleRepository) *UserService {
	return &UserService{
		repository:     r,
		roleRepository: roleRepository,
	}
}

//SignupService is used for creating new user
func (service *UserService) SignupService(userInfo types.SignupRequest) error {

	//Check email is already used
	_, err := service.repository.FindByEmail(userInfo.Email)

	if err == nil {
		return errors.New("The email is already used for another account!")
	}

	user := NewUser(userInfo.Name, userInfo.Email, userInfo.Password, 1)

	//check validations
	err2 := user.Validate()

	if err2 != nil {
		return err2
	}

	//Add new user to database
	err3 := service.repository.Create(user)

	if err3 != nil {
		return err3
	}

	return nil
}

//SignInService is used for login the system, it checks user credentials
func (service *UserService) SignInService(signinInfo types.SigninRequest) (User, error) {

	//Check user info is exist
	user, err := service.repository.FindByEmail(signinInfo.Email)

	if err != nil {
		log.Println(err)
		return User{}, err
	}

	//check password
	data := []byte(signinInfo.Password)
	hash := sha256.Sum256(data)
	strHash := fmt.Sprintf("%x", hash)

	if strHash != user.PasswordHash {
		log.Println("User password is not matched!")
		return User{}, errors.New("")
	}

	return user, nil
}

//SignOutService is used for logout from system, it destroy token
func (service *UserService) SignOutService(token, email string, client redis.RedisClient) error {

	var cachedToken string

	err := client.GetKey(email, cachedToken)

	if err != nil {
		return err
	}

	if token != cachedToken {
		return ErrUserNotAuthorized
	}

	err1 := client.DeleteKey(email)

	if err1 != nil {
		return err1
	}

	return nil
}

func (service *UserService) InsertSampleData() {
	tableExist := service.repository.db.Migrator().HasTable(&User{})

	if !tableExist {
		service.repository.MigrateTable()

		//admin
		admin := NewUser("admin", "admin@picus.com", "admin", 2)
		service.repository.Create(admin)

		//user
		user := NewUser("user", "user@picus.com", "1234", 1)
		service.repository.Create(user)

	}
}
