package server

import (
	"log"
	"net/http"
)

type server struct {
	port     string
	handlers []http.HandleFunc
}

// start intiates the http server in the given port
func (serv *server) start() error {
	log.Printf("Starting server at port %s", server.port)
	return http.ListenAndServe(serv.port, nil)
}

// registerHandlers registers all handlers to default mux server
func (serv *server) registerHandlers() error {
	log.Println("Registering all handlers to default mux server")
	serv.handlers = []http.HandleFunc{
		http.HandleFunc("/getFeeds/", getFeeds),
		http.HandleFunc("/getFollowers/", getFollowers),
		http.HandleFunc("/getFollowees/", getFollowees),
		http.HandleFunc("/addPost/", addPost),
		http.HandleFunc("/deletePost/", deletePost),
		http.HandleFunc("/addComment/", addComment),
		http.HandleFunc("/deleteComment/", deleteComment),
		http.HandleFunc("/follow/", follow),
		http.HandleFunc("/unFollow/", unFollow),
	}

	return nil
}

// parseServerConf parses the server.yaml file
func parseServerConf() (*server, error) {
	log.Println("Parsing server.yaml file")
	serverConfMap, err := conf.ParseConfFile("./../conf/server.yaml")
	if err != nil {
		return nil, err
	}

	return &server{port: serverConfMap["port"]}, nil
}

// StartServer will start the server
func StartServer() error {

	serv, err := parseServerConf()
	if err != nil {
		return err
	}

	err = serv.registerHandlers()
	if err != nil {
		return err
	}

	err = serv.start()
	if err != nil {
		return err
	}
}
