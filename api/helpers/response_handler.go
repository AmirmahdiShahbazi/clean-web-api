package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code int `json:"code"`
}

type ValidationErrorResponse struct {
	Response
	Errors map[string]string `json:"errors"`
}

type SuccessfulResponse struct {
	Response
	Data map[string]any `json:"data"`
}

type ErrorResponse struct {
	Response
	Error string `json:"error"`
}
var errorResponse ErrorResponse
func MakeValidationErrors(err error) *ValidationErrorResponse {
	var ve validator.ValidationErrors
	if err == nil {
		return &ValidationErrorResponse{}
	}
	if errors.As(err, &ve) {
		out := make(map[string]string, len(ve))
		for _, fe := range ve {
			out[fe.Field()] = fe.Tag()
		}
		return &ValidationErrorResponse{
			Response: Response{
				Code: 422,
			},
			Errors: out,
		}
	}
	return &ValidationErrorResponse{}
}

func HandleValidationErrors(c *gin.Context, validator interface{}) interface{} {
	c.Writer.Header().Set("Content-Type", "application/json")
	if err := c.Bind(&validator); err != nil {
		validationErrors := MakeValidationErrors(err)
		if validationErrors.Errors != nil {
			c.JSON(http.StatusBadRequest, validationErrors)
			return nil
		}
	}
	return &validator
}

func HandleSuccessfulResponse(c *gin.Context, data map[string]any) {
	var responseHandler SuccessfulResponse
	responseHandler.Code = http.StatusOK
	responseHandler.Data = data
	c.JSON(http.StatusOK, responseHandler)
}

func HandleErrorResponse(c *gin.Context, code int, error string){
	var responseHandler ErrorResponse
	responseHandler.Code = code
	responseHandler.Error = error
	c.AbortWithStatusJSON(code, responseHandler)
}
