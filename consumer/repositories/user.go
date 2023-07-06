package repositories

import (
	repo "events/repositories"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(repo.User) error
	FindById(id string) (repo.User, error)
	DeleteUser(id string) error
	GetUserPassword(email string) (string, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.Table("user").AutoMigrate(&repo.User{})
	return userRepository{db: db}
}

func (u userRepository) CreateUser(userData repo.User) error {
	return u.db.Table("user").Save(userData).Error
}

func (u userRepository) FindById(userId string) (repo.User, error) {
	var user repo.User
	err := u.db.Table("user").Where("id = ?", userId).First(&user).Error
	return user, err
}

func (u userRepository) DeleteUser(userId string) error {
	return u.db.Table("user").Where("id = ?", userId).Delete(&repo.User{}).Error
}

func (u userRepository) GetUserPassword(email string) (string, error) {
	var user repo.User
	err := u.db.Table("user").Where("email = ?", email).Select("password").First(&user).Error
	return user.Password, err
}
