package service

import (
	"github.com/boldlogic/cbr-market-data-worker/internal/client"
	"github.com/boldlogic/cbr-market-data-worker/internal/models"
	"github.com/boldlogic/cbr-market-data-worker/internal/service/request_catalog"
	"github.com/sirupsen/logrus"
)

type Service struct {
	client   *client.Client
	Provider *request_catalog.Provider
	//Storage      *storage.Storage
	log          logrus.FieldLogger
	CurrencyRepo CurrencyRepository
	fxRateRepo   FxRateRepository
}

func NewService(cl *client.Client, registry *request_catalog.Provider, currencyRepo CurrencyRepository, fxRateRepo FxRateRepository, log logrus.FieldLogger) *Service {

	return &Service{
		client:   cl,
		Provider: registry,
		//Storage:      storage,
		CurrencyRepo: currencyRepo,
		fxRateRepo:   fxRateRepo,
		log:          log,
	}
}

type CurrencyRepository interface {
	SaveCurrencies([]models.Currency) []error
	GetCurrencies() ([]models.Currency, error)
	GetCurrency(charCode string) (models.Currency, error)
}

type FxRateRepository interface {
	SaveFxRates([]models.FxRate) []error
}
