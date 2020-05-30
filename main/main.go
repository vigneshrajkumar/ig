package main

import "github.com/vigneshrajkumar/ig/server"

// start of the program
func main() {
	err := server.StartServer()
	if err != nil {
		log.Errorln("Error due to starting the server", err)
		return
	}
}
