package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Algorithm_3(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented",
	})
}
