package service

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/boldlogic/cbr-market-data-worker/internal/models"
	"github.com/boldlogic/cbr-market-data-worker/pkg/xmlconv"
	"golang.org/x/net/html/charset"
)

type Valute struct {
	ID         string          `xml:"ID,attr"`
	ISONumCode string          `xml:"NumCode"`
	CharCode   string          `xml:"CharCode"`
	Nominal    int             `xml:"Nominal"`
	Name       string          `xml:"Name"`
	Value      xmlconv.RuFloat `xml:"Value"`
	VunitRate  xmlconv.RuFloat `xml:"VunitRate"`
}

type ValCurs struct {
	Date   string   `xml:"Date,attr"`
	Name   string   `xml:"name,attr"`
	Valute []Valute `xml:"Valute"`
}

func (c *Client) GetCurrencyRates(ctx context.Context, req *http.Request) error {
	resp, err := c.sendRequest(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Не удалось получить курсы валют: %d", resp.StatusCode)
	}
	decoder := xml.NewDecoder(bytes.NewReader(resp.Body))
	decoder.CharsetReader = charset.NewReaderLabel
	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return fmt.Errorf("Не удалось получить курсы валют: %w", err)
	}
	c.log.Infof("%s", valCurs)

	rateDate, err := time.Parse("02.01.2006", valCurs.Date)
	if err != nil {
		return fmt.Errorf("не удалось определить дату курсов валют")
	}

	rates := make([]models.FxRate, 0, len(valCurs.Valute))

	for _, item := range valCurs.Valute {

		isoCode := 0
		if item.ISONumCode != "" {
			parsed, err := strconv.Atoi(item.ISONumCode)
			if err == nil {
				isoCode = parsed
			}
		}
		if isoCode <= 0 {
			continue //наверное надо будет потом логировать кривые валюты, но пока не критично
		}
		rates = append(rates, models.FxRate{
			Date:             rateDate,
			QuoteISOCode:     643,
			BaseISOCode:      isoCode,
			Nominal:          item.Nominal,
			QuoteForNominal:  float64(item.Value),
			QuotePerUnit:     float64(item.VunitRate),
			BasePerQuoteUnit: float64(1 / item.VunitRate),
		})
	}
	errs := c.Storage.SaveFxRates(rates)
	if len(errs) > 0 {
		if err := errors.Join(errs...); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}
