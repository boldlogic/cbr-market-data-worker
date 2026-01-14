package v1

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Mux     *chi.Mux
	Handler *Handler
	logger  logrus.FieldLogger
	//config  *config.Config
}

func NewRouter(handler *Handler, log logrus.FieldLogger) *Router {
	r := chi.NewRouter()
	r.Post("/requests", handler.execRequest)
	return &Router{
		Mux:     r,
		Handler: handler,
		logger:  log,
		//config:  cfg,
	}
}
