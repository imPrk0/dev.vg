package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/imPrk0/dev.vg/pkg/conf"
)

func Init() {
	InitApplication()
	conf.InitConfig()

	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
