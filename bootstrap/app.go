package bootstrap

import (
	"fmt"
	"github.com/imPrk0/dev.vg/pkg/conf"
)

func InitApplication() {
	fmt.Print(`
   ██████╗ ███████╗██╗   ██╗     ██╗   ██╗ ██████╗ 
   ██╔══██╗██╔════╝██║   ██║     ██║   ██║██╔════╝ 
   ██║  ██║█████╗  ██║   ██║     ██║   ██║██║  ███╗
   ██║  ██║██╔══╝  ╚██╗ ██╔╝     ╚██╗ ██╔╝██║   ██║
   ██████╔╝███████╗ ╚████╔╝  ██╗  ╚████╔╝ ╚██████╔╝
   ╚═════╝ ╚══════╝  ╚═══╝   ╚═╝   ╚═══╝   ╚═════╝ 
	Version: ` + conf.Version + ` Author: Prk (x.com/imPrk_)
	Project: https://github.com/imPrk0/dev.vg
============================================================
`)
	go CheckUpdate()
}

func CheckUpdate() {
	// https://api.github.com/repos/imPrk0/dev.vg/releases
}
