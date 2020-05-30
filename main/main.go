package main

import (
	"log"
	"net/http"
)

// read requests
func getFeeds(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("getFeeds"))
}

func getFollowers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("getFollowers"))
}

func getFollowees(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("getFollowees"))
}

// write requests
func addPost(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("addPost"))
}

func deletePost(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("deletePost"))
}

func addComment(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("addComment"))
}

func deleteComment(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("deleteComment"))
}

func follow(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("follow"))
}

func unFollow(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("unFollow"))
}

func main() {
	// register all handlers to default mux server
	http.HandleFunc("/getFeeds/", getFeeds)
	http.HandleFunc("/getFollowers/", getFollowers)
	http.HandleFunc("/getFollowees/", getFollowees)
	http.HandleFunc("/addPost/", addPost)
	http.HandleFunc("/deletePost/", deletePost)
	http.HandleFunc("/addComment/", addComment)
	http.HandleFunc("/deleteComment/", deleteComment)
	http.HandleFunc("/follow/", follow)
	http.HandleFunc("/unFollow/", unFollow)

	log.Print("Starting Server at port 8080")
	// start the http server in default port
	log.Fatal(http.ListenAndServe(":8080", nil))
}
