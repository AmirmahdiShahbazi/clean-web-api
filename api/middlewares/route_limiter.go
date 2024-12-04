package middlewares

import (
	"net/http"

	"github.com/AmirmahdiShahbazi/clean-web-api/api/helpers"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)



func RateLimiter() gin.HandlerFunc {
	limiter := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
		if err != nil {
			helpers.HandleErrorResponse(c, http.StatusTooManyRequests, err.Error())
		} else {
			c.Next()
		}

	}
}
