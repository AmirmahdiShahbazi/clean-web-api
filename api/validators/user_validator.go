package validators



type CreateUserValidator struct{
	BaseValidator
	Name string `json:"name" binding:"required"`
}