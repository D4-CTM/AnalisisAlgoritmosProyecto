package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ColorGraph(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented",
	})
}
