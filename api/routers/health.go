package routers

import (
	"github.com/AmirmahdiShahbazi/clean-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func HealthCheck(r *gin.RouterGroup){
	handler := handlers.NewHealth()
	r.GET("/", handler.HealthCheck)
}