package services

import (
	"errors"
	"log"
	"os"
	"time"
	"todo-list-api/models"
	"todo-list-api/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//var jwtSecret = []byte("your_secret_key") // default secret; override via env var

// AuthService handles authentication business logic.
type AuthService interface {
	Register(user *models.User) (string, error)
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService returns a new instance of AuthService.
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

// Register creates a user, hashes the password, and returns a JWT token.
func (s *authService) Register(user *models.User) (string, error) {
	// Check if the user already exists.
	existingUser, _ := s.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return "", errors.New("user already exists")
	}

	// Hash the password.
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashed)
	//log.Printf("Hashed password for user %s: %s", user.Email, user.Password)

	// Store the user in the database.
	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}

	// Generate a JWT token.
	return generateToken(user.ID.Hex())
}

// Login verifies the user credentials and returns a JWT token.
func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		log.Printf("Error finding user by email %s: %v", email, err)
		return "", errors.New("invalid email or password")
	}
	// Ensure that a user with the given email exists.
	if user == nil {
		log.Printf("User with email %s not found", email)
		return "", errors.New("invalid email or password")
	}
	log.Printf("Attempting password comparison for user %s", email)
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Password comparison failed for user %s: %v", email, err)
		return "", errors.New("invalid email or password")
	}
	log.Printf("User %s logged in successfully", email)
	return generateToken(user.ID.Hex())
}

// generateToken creates a JWT token that expires in 72 hours.
func generateToken(userID string) (string, error) {
	// Use secret from environment if present.
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your_secret_key"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secret))
}
