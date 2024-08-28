package services

import (
	"github.com/chigaji/maron_speak/internal/domain/models"
	"github.com/chigaji/maron_speak/internal/repository/db"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *db.UserRepo
}

func NewAuthService(userRep *db.UserRepo) *AuthService {
	return &AuthService{UserRepo: userRep}
}

func (s *AuthService) Register(username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.UserRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	user, err := s.UserRepo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
