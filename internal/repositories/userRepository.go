package repositories

import (
	"fmt"
	"repeatro/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO: write repository part, choose bd(postgresql?)
type UserRepository struct {
	db *gorm.DB
}

func CreateNewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepositoryMock struct{}

type UserRepositoryInterface interface {
	CreateUser(user *models.User)  error
	ReadUser(user_id uuid.UUID) (*models.User, error)
	ReadAllUsers() ([]models.User, error)
	ReadUserByEmail(email string) (*models.User, error)
}

func (ur UserRepository) CreateUser(user *models.User) error {
	fmt.Println("TOO")
	return ur.db.Create(user).Error	
}

func (ur UserRepository) ReadUser(user_id uuid.UUID) (*models.User, error) {
	var user models.User
	err := ur.db.Where("user_id = ?", user_id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur UserRepository) ReadUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return &user, nil
}

func (ur UserRepository) ReadAllUsers() ([]models.User, error) {
	var users []models.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

