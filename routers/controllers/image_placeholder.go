package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleImage(c *gin.Context, param string) {
	c.Redirect(
		http.StatusFound,
		"https://sdfsdf.dev"+param,
	)
}
