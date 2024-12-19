package service

import (
	"errors"
	"time"

	"task-management-system/configs"
	"task-management-system/internal/models"
	"task-management-system/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	config   *configs.Config
}

func NewAuthService(userRepo *repository.UserRepository, config *configs.Config) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *AuthService) Register(signup *models.UserSignup) (*models.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(signup.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Username: signup.Username,
		Email:    signup.Email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(login *models.UserLogin) (string, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(login.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateJWTToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) generateJWTToken(user *models.User) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}