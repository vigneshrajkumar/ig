package server

import (
	"instagram/conf"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	port       string
	mainRouter *mux.Router
}

// start intiates the http server in the given port
func (serv *server) start() error {
	log.Printf("Starting server at port %s", serv.port)
	return http.ListenAndServe(serv.port, serv.mainRouter)
}

// registerHandlers registers all handlers to default mux server
func (serv *server) registerHandlers() error {
	log.Println("Registering all handlers to mux server")

	// declared main router instagram
	mainRouter := serv.mainRouter.PathPrefix("/instagram").Subrouter()

	// declared routes for users
	users := mainRouter.PathPrefix("/users").Subrouter()

	users.HandleFunc("/", createUser).
		Methods("Post", "Put")

	users.HandleFunc("/{user_id:[0-9]+}", deleteUser).
		Methods("Delete")

	users.HandleFunc("/", getAllUser).
		Methods("Get")

	// declared routes for unique user
	uniqueUser := users.PathPrefix("/{user_id:[0-9]+}").Subrouter()

	uniqueUser.HandleFunc("/posts", putUniqueUserPost).
		Methods("Post", "Put")

	uniqueUser.HandleFunc("/posts/{post_id:[0-9]+}", deleteUniqueUserPost).
		Methods("Delete")

	uniqueUser.HandleFunc("/posts", getUniqueUserPost).
		Methods("Get")

	uniqueUser.HandleFunc("/profile", getUniqueUserProfile).
		Methods("Get")

	uniqueUser.HandleFunc("/follow/{follow_id:[0-9]+", followUser).
		Methods("Post", "Put")

	uniqueUser.HandleFunc("/unfollow/{follow_id:[0-9]+", unfollowUser).
		Methods("Delete")

	uniqueUser.HandleFunc("/followers", getUniqueUserFollowers).
		Methods("Get")

	uniqueUser.HandleFunc("/followings", getUniqueUserFollowings).
		Methods("Get")

	uniqueUser.HandleFunc("/feeds", getUniqueUserFeeds).
		Methods("Get")

	// declared routes for unique post
	uniquePost := uniqueUser.PathPrefix("/posts/{post_id:[0-9]+}").Subrouter()

	uniquePost.HandleFunc("/comments", putUniquePostComment).
		Methods("Post", "Put")

	uniquePost.HandleFunc("/comments/{comment_id:[0-9]+}", deleteUniquePostComment).
		Methods("Delete")

	uniquePost.HandleFunc("/comments", getUniquePostComment).
		Methods("Get")

	return nil
}

// parseServerConf parses the server.yaml file
func parseServerConf() (*server, error) {
	log.Println("Parsing server.yaml file")
	serverConfMap, err := conf.ParseConfFile("/conf/server.yaml")
	if err != nil {
		return nil, err
	}

	mainRouter := mux.NewRouter()

	return &server{port: serverConfMap["port"], mainRouter: mainRouter}, nil
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

	return nil
}
