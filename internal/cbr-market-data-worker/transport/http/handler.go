package httpserver

import (
	"github.com/boldlogic/PortfolioLens/internal/cbr-market-data-worker/service"
	v1 "github.com/boldlogic/PortfolioLens/internal/cbr-market-data-worker/transport/http/v1"
	"github.com/sirupsen/logrus"
)

type Handler = v1.Handler

func NewHandler(logger logrus.FieldLogger, svc service.Service) *v1.Handler {
	return v1.NewHandler(logger, svc)
}
