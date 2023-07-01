package controller

import "net/http"

type Router interface {
	HandleTitlesRequest(w http.ResponseWriter, r *http.Request)
	HandlePagesRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc TitleController
	pc PageController
}

func NewRouter(tc TitleController, cc PageController) Router {
	return &router{tc, cc}
}

func (ro *router) HandleTitlesRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.tc.GetTitles(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ro *router) HandlePagesRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.pc.GetPages(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
