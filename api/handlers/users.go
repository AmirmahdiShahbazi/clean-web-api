package handlers

import (
	"github.com/AmirmahdiShahbazi/clean-web-api/api/helpers"
	"github.com/AmirmahdiShahbazi/clean-web-api/api/validators"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Create(c *gin.Context) {
	user := validators.CreateUserValidator{}
	if validatedUser := helpers.HandleValidationErrors(c, &user); validatedUser == nil {
		return
	}
	helpers.HandleSuccessfulResponse(c, map[string]any{"user": user})
}
