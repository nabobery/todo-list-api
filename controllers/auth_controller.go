package controllers

import (
	"net/http"
	"todo-list-api/models"
	"todo-list-api/services"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication related endpoints.
type AuthController struct {
	authService services.AuthService
}

// NewAuthController creates a new AuthController instance.
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

// Register handles user registration.
//
// @Summary Register a new user
// @Description Register a new user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 200 {object} map[string]string "token"
// @Router /register [post]
func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := ac.authService.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Login handles user authentication.
//
// @Summary Login user
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "User credentials"
// @Success 200 {object} map[string]string "token"
// @Router /login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := ac.authService.Login(creds.Email, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
