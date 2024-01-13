package controllers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func HandleRandomPick(c *gin.Context, param string) {
	matches := regexp.MustCompile(`^/random,(.*)$`).FindStringSubmatch(param)

	if 2 == len(matches) {
		splitStrings := strings.Split(matches[1], ",")
		rand.Seed(time.Now().UnixNano())
		c.JSON(http.StatusOK, gin.H{
			"data": splitStrings[rand.Intn(len(splitStrings))],
		})
	} else {
		NotFound404(c)
	}
}
