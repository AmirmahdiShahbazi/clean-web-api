package handlers

import (
	"github.com/AmirmahdiShahbazi/clean-web-api/api/helpers"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	helpers.HandleSuccessfulResponse(c, map[string]any{"message": "Working!!!"})
}
