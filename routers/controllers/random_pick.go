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
	matches := regexp.MustCompile(`^random,(.*)$`).FindStringSubmatch(param)
	if 2 == len(matches) {
		randomString := matches[1]
		splitStrings := strings.Split(randomString, ",")
		// 从切片中随机选择一个值
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(splitStrings))
		randomValue := splitStrings[randomIndex]
		c.String(http.StatusOK, "Matched Case 3: Random Value=%s", randomValue)
	} else {
		c.String(http.StatusBadRequest, "Invalid format for Case 3")
	}
}
