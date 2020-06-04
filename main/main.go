package main

import (
	"instagram/server"

	"github.com/golang/glog"
)

// start of the program
func main() {
	err := server.StartServer()
	if err != nil {
		glog.Errorln("Error due to starting the server", err)
		return
	}
}
