package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/imPrk0/dev.vg/pkg/util"
	"github.com/imPrk0/dev.vg/routers/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gzip.Gzip(
		gzip.DefaultCompression,
		gzip.WithExcludedPaths([]string{
			"/api/",
		}),
	))

	r.GET("/*param", func(c *gin.Context) {
		param := c.Param("param")

		switch {
		case util.MatchAndHandle(param, `^/(\d+)[swpt],?$`):
			controllers.HandleRandomText(c, param)
		case util.MatchAndHandle(param, `^/(\d+)x(\d+)\.(png|jpg),([a-zA-Z]+),([a-zA-Z]+),?$`):
			controllers.HandleImage(c, param)
		case util.MatchAndHandle(param, `^/random,(.*)$`):
			controllers.HandleRandomPick(c, param)
		case util.MatchAndHandle(param, `^/[a-zA-Z0-9]{6,16}$`):
			controllers.HandleShortLink(c, param)
		default:
			controllers.NotFound404(c)
		}
	})

	//api := r.Group("/api")

	/*
		Middleware
		Author: Prk
	*/

	return r
}

func InitCORS(router *gin.Engine) {
	//router.Use(cors.New(cors.Config{}))
}
