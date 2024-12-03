package handlers

import (
	"net/http"
	"github.com/AmirmahdiShahbazi/clean-web-api/api/helpers"
	"github.com/AmirmahdiShahbazi/clean-web-api/api/validators"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}


func NewUserHandler() *UserHandler{
	return &UserHandler{}
}

func (u *UserHandler) Create(c *gin.Context) {
	var user validators.CreateUserValidator
	if validatedUser := helpers.HandleValidationErrors(c, &user); validatedUser == nil {  
		return 
	}  	
	c.JSON(http.StatusOK, gin.H{"code":200 ,"data": user})
}
