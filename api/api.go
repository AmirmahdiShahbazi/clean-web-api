package api

import (
	"github.com/AmirmahdiShahbazi/clean-web-api/api/middlewares"
	"github.com/AmirmahdiShahbazi/clean-web-api/api/routers"
	"github.com/gin-gonic/gin"
)

func InitRouters() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middlewares.RateLimiter())

	v1 := router.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.HealthCheck(health)

		user:= v1.Group("/users")
		routers.User(user)
	}
	router.Run(":5000")
}
