package db

import (
	"github.com/chigaji/maron_speak/internal/domain/models"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewuserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepo) FindUserByUsername(username string) (*models.User, error) {

	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
