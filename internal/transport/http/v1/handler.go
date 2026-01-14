package v1

import "github.com/sirupsen/logrus"

type Handler struct {
	log logrus.FieldLogger
}

func NewHandler(logger logrus.FieldLogger) *Handler {
	return &Handler{
		log: logger,
	}
}
