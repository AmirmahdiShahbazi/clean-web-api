package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "Working!!")
}