package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imPrk0/dev.vg/pkg/util"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func HandleRandomText(c *gin.Context, param string) {
	matches := regexp.MustCompile(`^/(\d+)([swpt]),?$`).FindStringSubmatch(param)

	if 3 == len(matches) {
		var (
			number    int
			option    string
			isArray   bool
			data      []string
			apiData   map[string]any
			delimiter string
		)

		delimiter = " "
		number, _ = strconv.Atoi(matches[1])
		option = matches[2]
		isArray = strings.HasSuffix(param, ",")

		switch option {
		case "w":
			data = util.GenerateRandomWords(number)
		case "s":
			data = util.GenerateRandomSentence(number)
		case "p":
			data = util.GenerateRandomParagraph(number)
		case "t":
			data = util.GenerateRandomText(number)
		}

		if isArray && "t" != option {
			apiData = gin.H{
				"data": data,
			}
		} else {
			apiData = gin.H{
				"data": strings.Join(
					data,
					delimiter,
				),
			}
		}

		c.JSON(http.StatusOK, apiData)
	} else {
		NotFound404(c)
	}
}
