package factories

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessSingle[T any](c *gin.Context, message string, data T) {
	response := CreateSingleResponse(
		http.StatusOK,
		message,
		data,
	)
	c.JSON(http.StatusOK, response)
}

func Created[T any](c *gin.Context, message string, data T) {
	response := CreateSingleResponse(
		http.StatusCreated,
		message,
		data,
	)
	c.JSON(http.StatusCreated, response)
}

func SuccessList[T any](c *gin.Context, message string, data []T) {
	response := CreateListResponse(
		http.StatusOK,
		message,
		data,
	)
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, statusCode int, message string, errorMsg string, details any) {
	response := CreateSingleResponse(
		statusCode,
		message,
		errorMsg,
	)
	c.JSON(statusCode, response)
}

func NotFound(c *gin.Context, resource string) {
	Error(
		c,
		http.StatusNotFound,
		"Resource not found",
		resource+" not found",
		nil,
	)
}

func BadRequest(c *gin.Context, details any) {
	Error(
		c,
		http.StatusBadRequest,
		"Invalid request",
		"request validation failed",
		details,
	)
}

func InternalServerError(c *gin.Context, err error) {
	Error(
		c,
		http.StatusInternalServerError,
		"Internal server error",
		err.Error(),
		nil,
	)
}

func Unauthorized(c *gin.Context) {
	Error(
		c,
		http.StatusUnauthorized,
		"Unauthorized",
		"authentication required",
		nil,
	)
}
