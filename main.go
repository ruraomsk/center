package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/anoshenko/rui"
	"github.com/ruraomsk/ag-server/logger"
	"github.com/ruraomsk/scenter/snmps"
	"github.com/ruraomsk/scenter/web"
)

func main() {
	_ = os.MkdirAll("log/", 0777)
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := logger.Init("log/"); err != nil {
		log.Panic("Error logger system", err.Error())
		return
	}
	logger.Info.Print("Starting Server")
	go snmps.Start("0.0.0.0:161")
	go rui.AddEmbedResources(&resources)
	go web.Web()

	for {
		time.Sleep(time.Minute)
	}
}
