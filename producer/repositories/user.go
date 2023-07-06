package repositories

import (
	repo "events/repositories"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id string) (repo.User, error)
	FindByEmail(email string) (repo.User, error)
	FindIdCurrentUser(userId string) (user repo.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.Table("user").AutoMigrate(&repo.User{})
	return userRepository{db: db}
}
func (r userRepository) FindById(id string) (user repo.User, err error) {
	err = r.db.Table("user").
		Where("id=?", id).
		First(&user).
		Error
	return user, err
}

func (r userRepository) FindByEmail(email string) (user repo.User, err error) {
	err = r.db.
		Select("password, id, username, role").
		Where("email = ?", email).
		First(&user).
		Error
	return user, err
}

func (r userRepository) FindIdCurrentUser(userId string) (user repo.User, err error) {
	err = r.db.
		Select("password, id, username, role").
		Where("id = ?", userId).
		First(&user).
		Error
	return user, err
}
