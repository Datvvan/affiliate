package util

import "github.com/gin-gonic/gin"

func ErrorResponse(code int, message string) gin.H {
	return gin.H{
		"success": false,
		"code":    code,
		"error":   message,
	}
}

func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"success": true,
		"code":    200,
		"data":    data,
	}
}
