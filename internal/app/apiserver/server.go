package apiserver

import (
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"net/http"
	"rtlabs/internal/app/apiserver/apierror"
	"time"
)

type server struct {
	router *muxRouter
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)

}

type handlers struct {
	vote *voteHandle
}

type muxRouter struct {
	*http.ServeMux
	handler    *handlers
	middleware *Middleware
	apiUrl     *apiUrl
}

// Конфигурация роутера
func (mr *muxRouter) configureRouter() {
	url := mr.apiUrl
	mr.Handle(url.vote.base, mr.voteRouter())
}

type handleResponse struct {
	logger *zap.Logger
	//errors *store.ErrorStore
}

func (s *handleResponse) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	defer s.logger.Sync()
	s.logger.Info(err.Error()) //zap.String("url", r.URL.Path),
	zap.Int("status", code)
	zap.String("host", r.Host)
	zap.String("method", r.Method)
	zap.Duration("backoff", time.Since(time.Now()))
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *handleResponse) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer s.logger.Sync()
	s.logger.Info("new request",
		zap.String("method", r.Method),
		zap.Int("status", code),
		zap.String("url", r.URL.Path),
		zap.String("host", r.Host),
		zap.Duration("backoff", time.Since(time.Now())),
	)
}

func newServer(client Client) *server {
	logger, _ := zap.NewProduction()
	defer logger.Sync().Error()
	hr := &handleResponse{
		logger: logger,
		//errors: store.Errors,
	}
	h := &handlers{
		vote: &voteHandle{
			handleResponse: hr,
			store:          client,
		},
	}
	mwr := &Middleware{
		corsMiddleware:   corsMiddleware{},
		ErrorsMiddleware: apierror.ErrorsMiddleware{},
	}
	mr := &muxRouter{
		ServeMux:   http.NewServeMux(),
		handler:    h,
		middleware: mwr,
		apiUrl:     url,
	}

	s := &server{
		router: mr,
	}
	s.router.configureRouter()
	return s
}
