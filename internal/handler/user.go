package handler

import (
	"net/http"

	"github.com/backrock/saas-backend/internal/model"
	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	// TODO: Add service dependencies
}

// NewUserHandler creates a new user handler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Register creates a new user account
func (h *UserHandler) Register(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: Implement user registration logic
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registration endpoint - TODO",
	})
}

// Login authenticates a user and returns a JWT token
func (h *UserHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: Implement login logic
	c.JSON(http.StatusOK, gin.H{
		"message": "User login endpoint - TODO",
	})
}

// GetProfile returns the current user's profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	// TODO: Implement get profile logic
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user profile endpoint - TODO",
	})
}
