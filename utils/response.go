package utils

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, code int, message string, result interface{}) {
	response := gin.H{
		"code":    code,
		"message": message,
	}

	if result != nil {
		response["result"] = result
	}

	c.JSON(code, response)
}
