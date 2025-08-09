package services

import (
	"errors"
	"movie-tracker/database"
	"movie-tracker/models"
	"regexp"

	"gorm.io/gorm"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(username, email, password string) (*models.User, error) {
	if err := s.validateRegistration(username, email, password); err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
	}

	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	db := database.GetDB()
	if err := db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("username or email already exists")
		}
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password are required")
	}

	db := database.GetDB()
	var user models.User
	
	err := db.Where("username = ? OR email = ?", username, username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (s *AuthService) GetUserByID(id uint) (*models.User, error) {
	db := database.GetDB()
	var user models.User
	
	err := db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) validateRegistration(username, email, password string) error {
	if len(username) < 3 || len(username) > 50 {
		return errors.New("username must be between 3 and 50 characters")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}