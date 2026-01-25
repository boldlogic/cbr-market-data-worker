package service

import (
	"context"
	"fmt"
	"net/http"

	//	v1 "github.com/boldlogic/cbr-market-data-worker/internal/transport/http/v1"

	"github.com/boldlogic/cbr-market-data-worker/internal/client"
	"github.com/boldlogic/cbr-market-data-worker/internal/models"
)

func (c *Service) ExecuteTask(ctx context.Context, tsk models.Task) error {

	plan, err := c.Provider.GetPlan(tsk.Type)
	if err != nil {
		return err
	}
	var ccy models.Currency
	var reqParams = make(map[string]string)
	for _, param := range plan.Params {
		if param.Type == "cbcode" {
			ccy, err = c.CurrencyRepo.GetCurrency(tsk.Params[param.Type])

			if err != nil {
				return err
			}
			if ccy.CbCode == "" {
				c.log.Errorf("Не удалось определить цб код")

				return fmt.Errorf("Не удалось определить цб код")
			}
			reqParams[param.Name] = ccy.CbCode
			c.log.Infof("Определен Код ЦБ %s по Коду валюты %s", ccy.CbCode, tsk.Params[param.Type])
		}
		if tsk.Params[param.Type] != "" && param.Type != "cbcode" {
			reqParams[param.Name] = tsk.Params[string(param.Type)]
		}

	}

	req, err := c.client.PrepareRequestWithParams(ctx, plan, reqParams)
	if err != nil {
		return err
	}
	c.log.Infof("Подготовлен запрос к %s", req.URL)

	var resp client.Response
	cnt := 0
	for i := 0; i < plan.RetryCount+1; i++ {
		resp, err = c.client.SendRequest(ctx, req)

		if resp.StatusCode == http.StatusOK && err == nil {
			break
		}
		cnt++
	}
	if err != nil {
		c.log.Errorf("Ошибка при получении данных. Кол-во попыток: %v", cnt)
		return fmt.Errorf("Ошибка при получении данных")
	}
	if resp.StatusCode != http.StatusOK {
		c.log.Errorf("Запрос завершен c ошибкой. StatusCode: %d", resp.StatusCode)

		return fmt.Errorf("Запрос завершен c ошибкой. StatusCode: %d", resp.StatusCode)
	}
	c.log.Infof("Запрос завершен успешно. StatusCode: %d", resp.StatusCode)

	if tsk.Type == "CBR_CURRENCIES" {
		err = c.GetCbrCurrencies(ctx, resp.Body)
		if err != nil {
			return err
		}
	}
	if tsk.Type == "CBR_CURRENCY_RATES" {
		err = c.GetCurrencyRates(ctx, resp.Body)
		if err != nil {
			return err
		}
	}
	if tsk.Type == "CBR_CURRENCY_RATES_HISTORY" {
		err = c.GetCurrencyRatesDynamic(ctx, resp.Body, ccy)
		c.log.Infof("GetCurrencyRatesDynamic")

		if err != nil {
			return err
		}
	}
	c.log.Infof("Задание с типом %s выполнено успешно", tsk.Type)

	return nil

}
