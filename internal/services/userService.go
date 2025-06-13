package services

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"repeatro/internal/models"
	"repeatro/internal/repositories"
	"repeatro/internal/schemes"
	"repeatro/internal/security"
)

type UserService struct {
	userRepository repositories.UserRepositoryInterface
	security       *security.Security
}

type UserServiceMock struct{}

func CreateNewUserService(userRepository *repositories.UserRepository, security *security.Security) *UserService {
	return &UserService{
		userRepository: userRepository,
		security:       security,
	}
}

type UserServiceInterface interface {
	FindUser(user_id uuid.UUID) (*models.User, error)
	CreateUser(user *models.User) (uuid.UUID, error)
	FindAllUsers() ([]models.User, error)
	GetUserIdByEmail(email string) (uuid.UUID, error)
	GetUserByEmail(email string) (*models.User, error)
	Register(userRegister schemes.AuthUser) (string, error)
	Login(userLogin schemes.AuthUser) (string, error)
}

func (us *UserService) FindUser(user_id uuid.UUID) (*models.User, error) {
	return us.userRepository.ReadUser(user_id)
}

func (us *UserService) GetUserIdByEmail(email string) (uuid.UUID, error) {
	user, err := us.userRepository.ReadUserByEmail(email)
	if err != nil {
		return uuid.UUID{}, err
	}
	if reflect.DeepEqual(user, &models.User{}) {
		return uuid.UUID{}, fmt.Errorf("not found user")
	}
	return user.UserId, nil
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := us.userRepository.ReadUserByEmail(email)
	if err != nil {
		return &models.User{}, err
	}
	if reflect.DeepEqual(user, &models.User{}) {
		return &models.User{}, fmt.Errorf("not found user")
	}
	return user, nil
}

func (us *UserService) CreateUser(user *models.User) (uuid.UUID, error) {
	userInDB, err := us.FindUser(user.UserId)
	if err != nil {
		return uuid.UUID{}, err
	}

	if !reflect.DeepEqual(userInDB, &models.User{}) {
		return userInDB.UserId, nil
	}

	err = us.userRepository.CreateUser(user)
	if err != nil {
		return uuid.UUID{}, err
	}

	return user.UserId, nil
}

func (us *UserService) FindAllUsers() ([]models.User, error) {
	return us.userRepository.ReadAllUsers()
}

func (us *UserService) Register(userRegister schemes.AuthUser) (string, error) {
	userInDB, err := us.userRepository.ReadUserByEmail(userRegister.Email)
	if err != nil {
		if err.Error() != "record not found" {
			return "", err
		}
	}

	if userInDB != nil {
		return "", fmt.Errorf("cannot register user with same email twice")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost) // Use the default cost factor
	if err != nil {
		return "", err
	}

	user := models.User{
		Email:          userRegister.Email,
		HashedPassword: string(hashedPassword),
	}

	user_id, err := us.CreateUser(&user)
	if err != nil {
		return "", err
	}

	token, err := us.security.EncodeString(user.HashedPassword, user_id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *UserService) Login(userLogin schemes.AuthUser) (string, error) {
	// want to check that users exists and return user
	user, err := us.GetUserByEmail(userLogin.Email)
	if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userLogin.Password)) != nil {
		return "", err
	}

	// want to encode string and return token
	token, err := us.security.EncodeString(user.HashedPassword, user.UserId)
	if err != nil {
		return "", err
	}

	return token, nil
}
