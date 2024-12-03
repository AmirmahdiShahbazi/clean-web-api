package routers

import (
	"github.com/AmirmahdiShahbazi/clean-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup){
	handler := handlers.NewUserHandler()
	r.POST("/create", handler.Create)
}