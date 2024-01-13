package main

import (
	"github.com/imPrk0/dev.vg/routers"
	"net/http"
)

func init() {}

func main() {
	api := routers.InitRouter()
	server := &http.Server{
		Handler: api,
	}

	//signalChan := make(chan os.Signal, 1)
	//signal.Notify(
	//	signalChan,
	//	os.Interrupt,
	//	syscall.SIGTERM,
	//	syscall.SIGHUP,
	//	syscall.SIGQUIT,
	//)
	//go shutdown(signalChan, server)

	//defer func() {
	//	<-signalChan
	//}()

	server.Addr = ":8080"
	if err := server.ListenAndServe(); nil != err {
	}
}

//func shutdown(signalChan chan os.Signal, server *http.Server) {
//	//sig := <-signalChan
//	ctx := context.Background()
//
//	err := server.Shutdown(ctx)
//
//	if nil != err {
//	}
//
//	close(signalChan)
//}
