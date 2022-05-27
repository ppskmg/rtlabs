package apiserver

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rtlabs/internal/app/apiserver/apierror"
	"rtlabs/internal/app/model/vote"
	"rtlabs/internal/app/store"
	"rtlabs/pkg/mwr"
)

func (mr *muxRouter) voteRouter() *httprouter.Router {
	h := httprouter.New()
	h.POST(mr.apiUrl.vote.create,
		mwr.Middlewares(
			mr.middleware.cors,
		).Then(
			mr.handler.vote.create()))

	h.POST(mr.apiUrl.vote.delete,
		mwr.Middlewares(
			mr.middleware.cors,
		).Then(
			mr.handler.vote.delete()))

	h.GET(mr.apiUrl.vote.countVote,
		mwr.Middlewares(
			//mr.middleware.cors,
		).Then(
			mr.handler.vote.countVote()))
	h.POST(mr.apiUrl.vote.me,
		mwr.Middlewares(
			mr.middleware.cors,
		).Then(
			mr.handler.vote.myVote()))
	return h
}

type voteHandle struct {
	*handleResponse
	store Client
}

func (eh *voteHandle) create() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		req := &vote.Vote{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			eh.error(w, r, http.StatusBadRequest, err)
		}

		ee := &vote.Vote{
			WorkID:    req.WorkID,
			UserEmail: req.UserEmail,
			Contest:   req.Contest,
		}

		err := eh.store.Postgres.Vote().Create(ee)
		if err != nil {
			eh.error(w, r, http.StatusUnprocessableEntity, err)
		}
		eh.respond(w, r, http.StatusCreated, "vote +")
	}
}

func (eh *voteHandle) countVote() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var ee []vote.Count
		ee, err := eh.store.Postgres.Vote().ReadCountVote()
		if err != nil {
			if err.Error() == apierror.ErrNoRows.Message {
				eh.error(w, r, http.StatusNotFound, apierror.ErrRecordNotFound)
				return
			}
			eh.error(w, r, http.StatusBadRequest, err)
			return
		}
		eh.respond(w, r, http.StatusOK, ee)
	}
}

func (eh *voteHandle) myVote() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		req := &vote.Vote{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			eh.error(w, r, http.StatusBadRequest, err)
			return
		}

		var ee []vote.Vote
		ee, err := eh.store.Postgres.Vote().ReadMyVote(req.UserEmail, req.Contest)
		if err != nil {
			if err.Error() == apierror.ErrNoRows.Message {
				eh.error(w, r, http.StatusNotFound, apierror.ErrRecordNotFound)
				return
			}
			eh.error(w, r, http.StatusBadRequest, err)
			return
		}
		eh.respond(w, r, http.StatusOK, ee)
	}
}

func (eh *voteHandle) delete() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		req := &vote.Vote{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			eh.error(w, r, http.StatusBadRequest, err)
		}
		id, err := eh.store.Postgres.Vote().Delete(req.WorkID, req.UserEmail)

		if err != nil {
			if err == store.ErrRecordNotFound {
				eh.error(w, r, http.StatusNotFound, err)
			}
			eh.error(w, r, http.StatusUnprocessableEntity, err)
		}
		eh.respond(w, r, http.StatusOK, id)
	}
}
