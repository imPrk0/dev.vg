package controllers

import "github.com/gin-gonic/gin"

func NotFound404(c *gin.Context) {
	c.JSON(404, gin.H{
		"error": "Not Found",
	})
}
