package apiserver

import "github.com/julienschmidt/httprouter"

type Router interface {
	router() *httprouter.Router
}

type urlAuth struct {
	base     string
	login    string
	register string
}

type urlVote struct {
	base      string
	me        string
	create    string
	delete    string
	countVote string
}

type apiUrl struct {
	auth *urlAuth
	vote *urlVote
}

var (
	url = &apiUrl{
		vote: &urlVote{
			base:      "/api/vote/",
			me:        "/api/vote/me",
			create:    "/api/vote/create",
			delete:    "/api/vote/delete",
			countVote: "/api/vote/count",
		},
	}
)
