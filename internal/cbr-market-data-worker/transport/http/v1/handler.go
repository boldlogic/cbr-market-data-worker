package v1

import (
	"github.com/boldlogic/PortfolioLens/internal/cbr-market-data-worker/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Service service.Service
	log     logrus.FieldLogger
}

func NewHandler(logger logrus.FieldLogger, svc service.Service) *Handler {
	return &Handler{
		log:     logger,
		Service: svc,
	}
}
