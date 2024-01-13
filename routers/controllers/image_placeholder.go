package controllers

import (
	"github.com/gin-gonic/gin"
)

func HandleImage(c *gin.Context, param string) {
	c.JSON(200, gin.H{
		"data": c,
	})
}
