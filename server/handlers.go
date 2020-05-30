package server

import "net/http"

// Handlers

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
