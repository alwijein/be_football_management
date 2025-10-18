package utils

import (
	"net/http"

	"github.com/alwijein/be_football/internal/dto"
	"github.com/gin-gonic/gin"
)

// Meta represents API meta information
type Meta struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

// Response represents API response structure
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// SuccessResponse sends success response
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	var msg interface{} = message
	if message == "" {
		msg = nil
	}
	
	c.JSON(statusCode, Response{
		Meta: Meta{
			Code:    statusCode,
			Status:  "success",
			Message: msg,
		},
		Data: data,
	})
}

// SuccessResponseWithPagination sends success response with pagination
func SuccessResponseWithPagination(c *gin.Context, message string, data interface{}, pagination *dto.Pagination) {
	var msg interface{} = message
	if message == "" {
		msg = nil
	}
	
	responseData := gin.H{
		"items":      data,
		"pagination": pagination,
	}
	
	c.JSON(http.StatusOK, Response{
		Meta: Meta{
			Code:    http.StatusOK,
			Status:  "success",
			Message: msg,
		},
		Data: responseData,
	})
}

// ErrorResponse sends error response
func ErrorResponse(c *gin.Context, statusCode int, message string, errors interface{}) {
	var msg interface{} = message
	if message == "" {
		msg = nil
	}
	
	c.JSON(statusCode, Response{
		Meta: Meta{
			Code:    statusCode,
			Status:  "error",
			Message: msg,
		},
		Data: errors,
	})
}

// ValidationErrorResponse sends validation error response
func ValidationErrorResponse(c *gin.Context, errors interface{}) {
	ErrorResponse(c, http.StatusBadRequest, "Validation failed", errors)
}

// UnauthorizedResponse sends unauthorized response
func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

// NotFoundResponse sends not found response
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

// InternalServerErrorResponse sends internal server error response
func InternalServerErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message, nil)
}

// BadRequestResponse sends bad request response
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message, nil)
}

// ConflictResponse sends conflict response
func ConflictResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusConflict, message, nil)
}
