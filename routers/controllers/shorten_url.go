package controllers

import "github.com/gin-gonic/gin"

// HandleShortLink TODO
func HandleShortLink(c *gin.Context, param string) {
	c.JSON(200, gin.H{
		"data": param[1:],
	})
}
