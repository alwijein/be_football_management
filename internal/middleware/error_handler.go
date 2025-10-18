package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/alwijein/be_football/internal/utils"
)

// ValidationError represents validation error details
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorHandler handles errors and returns appropriate response
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Handle validation errors
			if validationErrors, ok := err.Err.(validator.ValidationErrors); ok {
				var errors []ValidationError
				for _, fieldError := range validationErrors {
					errors = append(errors, ValidationError{
						Field:   fieldError.Field(),
						Message: getValidationMessage(fieldError),
					})
				}
				utils.ValidationErrorResponse(c, errors)
				return
			}

			// Handle other errors
			utils.InternalServerErrorResponse(c, err.Error())
		}
	}
}

func getValidationMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fieldError.Field() + " is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fieldError.Field() + " must be at least " + fieldError.Param()
	case "max":
		return fieldError.Field() + " must be at most " + fieldError.Param()
	case "len":
		return fieldError.Field() + " must be " + fieldError.Param() + " characters"
	case "url":
		return "Invalid URL format"
	default:
		return fieldError.Field() + " is invalid"
	}
}
