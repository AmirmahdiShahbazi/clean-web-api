package helpers

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code int `json:"code"`
}

type ValidationErrorResponse struct {
	Response
	Errors map[string]string `json:"errors"`
}

func MakeValidationErrors(err error) *ValidationErrorResponse {
	var ve validator.ValidationErrors
	if(err == nil){
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
	if err := c.ShouldBindJSON(&validator); err != nil {
		validationErrors := MakeValidationErrors(err)  
		if validationErrors != nil {  
			c.JSON(http.StatusBadRequest, gin.H{"code": validationErrors.Code, "errors": validationErrors.Errors})  
			return nil 
		}
	}  
	return &validator  
} 