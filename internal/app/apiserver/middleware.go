package apiserver

import (
	"github.com/julienschmidt/httprouter"
	ae "rtlabs/internal/app/apiserver/apierror"

	"net/http"
)

type Middleware struct {
	corsMiddleware
	ae.ErrorsMiddleware
}

type corsMiddleware struct {
	store Client
	*handleResponse
}

func (c *corsMiddleware) cors(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "https://rtlabs.party")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next(w, r, params)
	}
}
